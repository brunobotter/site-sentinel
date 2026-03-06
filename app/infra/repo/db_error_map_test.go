package repo

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
)

func TestMapPostgresValidationError(t *testing.T) {
	validationMap := postgresValidationMap{
		ConstraintMessages: map[string]map[string]string{
			uniqueViolationCode: {
				"ux_tenants_name": "name já cadastrado",
			},
			checkViolationCode: {
				"chk_environment_name": "environment deve ser dev, stg ou prod",
			},
		},
		DefaultMessages: map[string]string{
			uniqueViolationCode:  "registro já cadastrado",
			checkViolationCode:   "valor inválido",
			notNullViolationCode: "campo obrigatório não informado",
			invalidTextCode:      "valor inválido para o tipo esperado",
		},
		NoRowsMessage: "registro não encontrado",
	}

	t.Run("returns not found for no rows", func(t *testing.T) {
		mappedErr := mapPostgresValidationError(sql.ErrNoRows, validationMap)
		if mappedErr == nil {
			t.Fatal("expected mapped error, got nil")
		}

		if mappedErr.Error() != "registro não encontrado" {
			t.Fatalf("expected 'registro não encontrado', got %q", mappedErr.Error())
		}
	})

	t.Run("returns specific message for mapped constraint", func(t *testing.T) {
		err := &pgconn.PgError{Code: uniqueViolationCode, ConstraintName: "ux_tenants_name"}

		mappedErr := mapPostgresValidationError(err, validationMap)
		if mappedErr == nil {
			t.Fatal("expected mapped error, got nil")
		}

		if mappedErr.Error() != "name já cadastrado" {
			t.Fatalf("expected 'name já cadastrado', got %q", mappedErr.Error())
		}
	})

	t.Run("returns default message when code is mapped but constraint is not", func(t *testing.T) {
		err := &pgconn.PgError{Code: uniqueViolationCode, ConstraintName: "ux_unknown"}

		mappedErr := mapPostgresValidationError(err, validationMap)
		if mappedErr == nil {
			t.Fatal("expected mapped error, got nil")
		}

		if mappedErr.Error() != "registro já cadastrado" {
			t.Fatalf("expected 'registro já cadastrado', got %q", mappedErr.Error())
		}
	})

	t.Run("returns default message when only SQLSTATE is mapped", func(t *testing.T) {
		err := &pgconn.PgError{Code: notNullViolationCode}

		mappedErr := mapPostgresValidationError(err, validationMap)
		if mappedErr == nil {
			t.Fatal("expected mapped error, got nil")
		}

		if mappedErr.Error() != "campo obrigatório não informado" {
			t.Fatalf("expected 'campo obrigatório não informado', got %q", mappedErr.Error())
		}
	})

	t.Run("returns nil for unmapped SQLSTATE", func(t *testing.T) {
		err := &pgconn.PgError{Code: "40001", ConstraintName: "ux_tenants_name"}

		mappedErr := mapPostgresValidationError(err, validationMap)
		if mappedErr != nil {
			t.Fatalf("expected nil, got %v", mappedErr)
		}
	})

	t.Run("returns nil for non database errors", func(t *testing.T) {
		mappedErr := mapPostgresValidationError(errors.New("generic error"), validationMap)
		if mappedErr != nil {
			t.Fatalf("expected nil, got %v", mappedErr)
		}
	})
}
