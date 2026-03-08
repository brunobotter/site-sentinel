package usecase

import (
	"context"

	"github.com/brunobotter/site-sentinel/application/domain"
	"github.com/brunobotter/site-sentinel/application/service"
)

type listTargetsUseCase struct {
	targetService service.TargetService
}

// NewListTargetsUseCase cria o caso de uso responsável por listar alvos.
func NewListTargetsUseCase(targetService service.TargetService) ListTargetsUseCase {
	return &listTargetsUseCase{targetService: targetService}
}

// Execute retorna os alvos vindo da camada de serviço.
func (u *listTargetsUseCase) Execute(ctx context.Context) ([]domain.MonitorTarget, error) {
	return u.targetService.List(ctx)
}
