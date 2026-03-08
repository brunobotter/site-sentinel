package repo

import (
	"context"
	"time"

	"github.com/brunobotter/site-sentinel/application/domain"
	"github.com/brunobotter/site-sentinel/infra/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type monitorTargetPgRepository struct {
	db  *pgxpool.Pool
	log logger.Logger
}

// NewMonitorTargetPgRepository cria o repositório PostgreSQL de monitor_targets.
func NewMonitorTargetPgRepository(db *pgxpool.Pool, log logger.Logger) *monitorTargetPgRepository {
	return &monitorTargetPgRepository{db: db, log: log}
}

// Create insere um novo alvo no banco.
func (r *monitorTargetPgRepository) Create(ctx context.Context, target domain.MonitorTarget) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO monitor_targets (
			id, name, url, method, timeout_ms, expected_status, retries, retry_delay_ms, active, created_at, updated_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
	`, target.ID, target.Name, target.URL, target.Method, target.Policy.Timeout.Milliseconds(), target.Policy.ExpectedStatus, target.Policy.Retries, target.Policy.RetryDelay.Milliseconds(), target.Active, target.CreatedAt, target.UpdatedAt)
	if mapped := mapPostgresValidationError(err, monitorTargetValidationMap); mapped != nil {
		return mapped
	}
	return err
}

// Update atualiza os dados mutáveis de um alvo existente.
func (r *monitorTargetPgRepository) Update(ctx context.Context, target domain.MonitorTarget) error {
	target.UpdatedAt = time.Now()
	_, err := r.db.Exec(ctx, `
		UPDATE monitor_targets
		SET name = $2, url = $3, method = $4, timeout_ms = $5, expected_status = $6, retries = $7, retry_delay_ms = $8, active = $9, updated_at = $10
		WHERE id = $1
	`, target.ID, target.Name, target.URL, target.Method, target.Policy.Timeout.Milliseconds(), target.Policy.ExpectedStatus, target.Policy.Retries, target.Policy.RetryDelay.Milliseconds(), target.Active, target.UpdatedAt)
	if mapped := mapPostgresValidationError(err, monitorTargetValidationMap); mapped != nil {
		return mapped
	}
	return err
}

// Delete remove um alvo por ID.
func (r *monitorTargetPgRepository) Delete(ctx context.Context, id string) error {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, `DELETE FROM monitor_targets WHERE id = $1`, uuidID)
	if mapped := mapPostgresValidationError(err, monitorTargetValidationMap); mapped != nil {
		return mapped
	}
	return err
}

// FindByID retorna um alvo específico por UUID.
func (r *monitorTargetPgRepository) FindByID(ctx context.Context, id string) (*domain.MonitorTarget, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(ctx, `
		SELECT id, name, url, method, timeout_ms, expected_status, retries, retry_delay_ms, active, created_at, updated_at
		FROM monitor_targets
		WHERE id = $1
	`, uuidID)

	target, err := scanMonitorTarget(row)
	if mapped := mapPostgresValidationError(err, monitorTargetValidationMap); mapped != nil {
		return nil, mapped
	}
	if err != nil {
		return nil, err
	}
	return target, nil
}

// List retorna alvos com paginação simples de limit/offset.
func (r *monitorTargetPgRepository) List(ctx context.Context, limit int, offset int) ([]domain.MonitorTarget, error) {
	if limit <= 0 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	rows, err := r.db.Query(ctx, `
		SELECT id, name, url, method, timeout_ms, expected_status, retries, retry_delay_ms, active, created_at, updated_at
		FROM monitor_targets
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return collectMonitorTargets(rows)
}

// ListActive retorna somente os alvos ativos para execução do scheduler.
func (r *monitorTargetPgRepository) ListActive(ctx context.Context) ([]domain.MonitorTarget, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, name, url, method, timeout_ms, expected_status, retries, retry_delay_ms, active, created_at, updated_at
		FROM monitor_targets
		WHERE active = true
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return collectMonitorTargets(rows)
}

type monitorTargetScanner interface {
	Scan(dest ...any) error
}

type monitorTargetRows interface {
	Next() bool
	Scan(dest ...any) error
	Err() error
}

func scanMonitorTarget(scanner monitorTargetScanner) (*domain.MonitorTarget, error) {
	var (
		target       domain.MonitorTarget
		timeoutMS    int64
		retryDelayMS int64
	)

	err := scanner.Scan(
		&target.ID,
		&target.Name,
		&target.URL,
		&target.Method,
		&timeoutMS,
		&target.Policy.ExpectedStatus,
		&target.Policy.Retries,
		&retryDelayMS,
		&target.Active,
		&target.CreatedAt,
		&target.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	target.Policy.Timeout = time.Duration(timeoutMS) * time.Millisecond
	target.Policy.RetryDelay = time.Duration(retryDelayMS) * time.Millisecond
	return &target, nil
}

// collectMonitorTargets percorre o cursor do banco e monta a lista final de entidades.
func collectMonitorTargets(rows monitorTargetRows) ([]domain.MonitorTarget, error) {
	targets := make([]domain.MonitorTarget, 0)
	for rows.Next() {
		target, err := scanMonitorTarget(rows)
		if err != nil {
			return nil, err
		}
		targets = append(targets, *target)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return targets, nil
}

var monitorTargetValidationMap = postgresValidationMap{
	ConstraintMessages: map[string]map[string]string{},
	DefaultMessages: map[string]string{
		notNullViolationCode: "campos obrigatorios de monitor target ausentes",
		invalidTextCode:      "dados invalidos para monitor target",
	},
	NoRowsMessage: "monitor target nao encontrado",
}
