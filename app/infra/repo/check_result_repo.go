package repo

import (
	"context"
	"time"

	"github.com/brunobotter/site-sentinel/application/domain"
	"github.com/brunobotter/site-sentinel/infra/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type checkResultPgRepository struct {
	db  *pgxpool.Pool
	log logger.Logger
}

// NewCheckResultPgRepository cria o repositório PostgreSQL de resultados de checks.
func NewCheckResultPgRepository(db *pgxpool.Pool, log logger.Logger) *checkResultPgRepository {
	return &checkResultPgRepository{db: db, log: log}
}

// Save persiste um único resultado.
func (r *checkResultPgRepository) Save(ctx context.Context, result domain.CheckResult) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO check_results (id, target_id, status_code, response_time_ms, is_up, error, checked_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
	`, result.ID, result.TargetID, result.StatusCode, result.ResponseTime.Milliseconds(), result.IsUp, result.Error, result.CheckedAt)
	if mapped := mapPostgresValidationError(err, checkResultValidationMap); mapped != nil {
		return mapped
	}
	return err
}

// SaveBatch persiste vários resultados na mesma transação para reduzir overhead.
func (r *checkResultPgRepository) SaveBatch(ctx context.Context, results []domain.CheckResult) error {
	if len(results) == 0 {
		return nil
	}

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, result := range results {
		_, err = tx.Exec(ctx, `
			INSERT INTO check_results (id, target_id, status_code, response_time_ms, is_up, error, checked_at)
			VALUES ($1,$2,$3,$4,$5,$6,$7)
		`, result.ID, result.TargetID, result.StatusCode, result.ResponseTime.Milliseconds(), result.IsUp, result.Error, result.CheckedAt)
		if mapped := mapPostgresValidationError(err, checkResultValidationMap); mapped != nil {
			return mapped
		}
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

// ListLatestByTarget busca histórico recente de um alvo específico.
func (r *checkResultPgRepository) ListLatestByTarget(
	ctx context.Context,
	targetID string,
	limit int,
) ([]domain.CheckResult, error) {
	uuidID, err := uuid.Parse(targetID)
	if err != nil {
		return nil, err
	}

	if limit <= 0 {
		limit = 50
	}

	rows, err := r.db.Query(ctx, `
		SELECT id, target_id, status_code, response_time_ms, is_up, error, checked_at
		FROM check_results
		WHERE target_id = $1
		ORDER BY checked_at DESC
		LIMIT $2
	`, uuidID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return collectCheckResults(rows)
}

// ListLatestGlobal busca os resultados mais recentes de todos os alvos.
func (r *checkResultPgRepository) ListLatestGlobal(
	ctx context.Context,
	limit int,
) ([]domain.CheckResult, error) {
	if limit <= 0 {
		limit = 50
	}

	rows, err := r.db.Query(ctx, `
		SELECT id, target_id, status_code, response_time_ms, is_up, error, checked_at
		FROM check_results
		ORDER BY checked_at DESC
		LIMIT $1
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return collectCheckResults(rows)
}

type checkResultScanner interface {
	Scan(dest ...any) error
}

type checkResultRows interface {
	Next() bool
	Scan(dest ...any) error
	Err() error
}

func scanCheckResult(scanner checkResultScanner) (*domain.CheckResult, error) {
	var (
		result         domain.CheckResult
		responseTimeMS int64
	)

	err := scanner.Scan(
		&result.ID,
		&result.TargetID,
		&result.StatusCode,
		&responseTimeMS,
		&result.IsUp,
		&result.Error,
		&result.CheckedAt,
	)
	if err != nil {
		return nil, err
	}

	result.ResponseTime = time.Duration(responseTimeMS) * time.Millisecond
	return &result, nil
}

// collectCheckResults converte as linhas retornadas pelo banco em entidades de domínio.
func collectCheckResults(rows checkResultRows) ([]domain.CheckResult, error) {
	results := make([]domain.CheckResult, 0)
	for rows.Next() {
		result, err := scanCheckResult(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, *result)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

var checkResultValidationMap = postgresValidationMap{
	ConstraintMessages: map[string]map[string]string{},
	DefaultMessages: map[string]string{
		notNullViolationCode:    "campos obrigatorios de check result ausentes",
		foreignKeyViolationCode: "target do check nao encontrado",
		invalidTextCode:         "dados invalidos para check result",
	},
	NoRowsMessage: "check result nao encontrado",
}
