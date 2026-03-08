package usecase

import (
	"context"

	"github.com/brunobotter/site-sentinel/application/command"
	"github.com/brunobotter/site-sentinel/application/service"
)

type createTargetUseCase struct {
	targetService service.TargetService
}

// NewCreateTargetUseCase monta o caso de uso de criação de alvo.
func NewCreateTargetUseCase(targetService service.TargetService) CreateTargetUseCase {
	return &createTargetUseCase{targetService: targetService}
}

// Execute delega para a camada de serviço, mantendo o use case fino e objetivo.
func (u *createTargetUseCase) Execute(ctx context.Context, cmd command.CreateTargetCommand) error {
	return u.targetService.Create(ctx, cmd)
}
