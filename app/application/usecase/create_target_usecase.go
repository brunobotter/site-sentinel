package usecase

import (
	"context"

	"github.com/brunobotter/site-sentinel/application/command"
	"github.com/brunobotter/site-sentinel/application/service"
)

type createTargetUseCase struct {
	targetService service.TargetService
}

func NewCreateTargetUseCase(targetService service.TargetService) CreateTargetUseCase {
	return &createTargetUseCase{targetService: targetService}
}

func (u *createTargetUseCase) Execute(ctx context.Context, cmd command.CreateTargetCommand) error {
	return u.targetService.Create(ctx, cmd)
}
