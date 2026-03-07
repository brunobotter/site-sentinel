package usecase

import (
	"context"

	"github.com/brunobotter/site-sentinel/application/command"
	"github.com/brunobotter/site-sentinel/application/service"
)

type runBatchCheckUseCase struct {
	checkExecutionService service.CheckExecutionService
}

func NewRunBatchCheckUseCase(
	checkExecutionService service.CheckExecutionService,
) RunBatchCheckUseCase {
	return &runBatchCheckUseCase{checkExecutionService: checkExecutionService}
}

func (u *runBatchCheckUseCase) Execute(ctx context.Context, cmd command.RunCheckBatchCommand) error {
	return u.checkExecutionService.RunBatch(ctx, cmd.Targets)
}
