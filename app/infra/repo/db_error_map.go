package repo

import (
	"database/sql"
	"errors"

	"github.com/brunobotter/site-sentinel/application"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	uniqueViolationCode     = "23505"
	foreignKeyViolationCode = "23503"
	checkViolationCode      = "23514"
	notNullViolationCode    = "23502"
	invalidTextCode         = "22P02"
)

type postgresValidationMap struct {
	ConstraintMessages map[string]map[string]string
	DefaultMessages    map[string]string
	NoRowsMessage      string
}

// mapPostgresValidationError traduz erros técnicos do Postgres para erros de aplicação.
//
// Para júnior: isso evita "vazar" mensagens de banco para a API e padroniza respostas.
func mapPostgresValidationError(err error, validationMap postgresValidationMap) error {
	if errors.Is(err, sql.ErrNoRows) || errors.Is(err, pgx.ErrNoRows) {
		return application.NewNotFoundApplicationError(application.ValidationDomain, errors.New(validationMap.NoRowsMessage))
	}

	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return nil
	}

	if codeConstraints, ok := validationMap.ConstraintMessages[pgErr.Code]; ok {
		if message, ok := codeConstraints[pgErr.ConstraintName]; ok {
			return application.NewValidationApplicationError(application.ValidationDomain, errors.New(message))
		}
	}

	if defaultMessage, ok := validationMap.DefaultMessages[pgErr.Code]; ok {
		return application.NewValidationApplicationError(application.ValidationDomain, errors.New(defaultMessage))
	}

	return nil
}
