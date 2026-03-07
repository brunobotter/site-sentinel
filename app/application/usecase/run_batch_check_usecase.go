package usecase

import (
	"context"

	"github.com/brunobotter/site-sentinel/application/domain"
	"github.com/brunobotter/site-sentinel/application/service"
)

type listLatestResultsUseCase struct {
	resultService service.CheckResultService
}

func NewListLatestResultsUseCase(resultService service.CheckResultService) ListLatestResultsUseCase {
	return &listLatestResultsUseCase{resultService: resultService}
}

func (u *listLatestResultsUseCase) Execute(ctx context.Context, limit int) ([]domain.CheckResult, error) {
	return u.resultService.ListLatestGlobal(ctx, limit)
}
