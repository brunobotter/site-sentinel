package usecase

import (
	"context"

	"github.com/brunobotter/site-sentinel/application/domain"
	"github.com/brunobotter/site-sentinel/application/service"
)

type listLatestResultsUseCase struct {
	resultService service.CheckResultService
}

// NewListLatestResultsUseCase cria o caso de uso de consulta para o dashboard.
func NewListLatestResultsUseCase(resultService service.CheckResultService) ListLatestResultsUseCase {
	return &listLatestResultsUseCase{resultService: resultService}
}

// Execute retorna os resultados mais recentes obedecendo ao limite solicitado.
func (u *listLatestResultsUseCase) Execute(ctx context.Context, limit int) ([]domain.CheckResult, error) {
	return u.resultService.ListLatestGlobal(ctx, limit)
}
