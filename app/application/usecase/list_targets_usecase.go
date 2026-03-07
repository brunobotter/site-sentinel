package usecase

import (
	"context"

	"github.com/brunobotter/site-sentinel/application/domain"
	"github.com/brunobotter/site-sentinel/application/service"
)

type listTargetsUseCase struct {
	targetService service.TargetService
}

func NewListTargetsUseCase(targetService service.TargetService) ListTargetsUseCase {
	return &listTargetsUseCase{targetService: targetService}
}

func (u *listTargetsUseCase) Execute(ctx context.Context) ([]domain.MonitorTarget, error) {
	return u.targetService.List(ctx)
}
