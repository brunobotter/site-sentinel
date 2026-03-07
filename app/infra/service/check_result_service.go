package service

import (
	"context"

	"github.com/brunobotter/site-sentinel/application/domain"
	"github.com/brunobotter/site-sentinel/application/repo"
)

type checkResultService struct {
	resultRepo repo.CheckResultRepository
}

func NewCheckResultService(resultRepo repo.CheckResultRepository) *checkResultService {
	return &checkResultService{resultRepo: resultRepo}
}

func (s *checkResultService) ListLatestGlobal(ctx context.Context, limit int) ([]domain.CheckResult, error) {
	if limit <= 0 {
		limit = 50
	}
	return s.resultRepo.ListLatestGlobal(ctx, limit)
}
