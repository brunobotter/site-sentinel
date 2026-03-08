package usecase

import (
	"context"

	"github.com/brunobotter/site-sentinel/application/command"
	"github.com/brunobotter/site-sentinel/application/service"
)

type runBatchCheckUseCase struct {
	checkExecutionService service.CheckExecutionService
}

// NewRunBatchCheckUseCase cria o caso de uso que dispara checks em lote.
func NewRunBatchCheckUseCase(
	checkExecutionService service.CheckExecutionService,
) RunBatchCheckUseCase {
	return &runBatchCheckUseCase{checkExecutionService: checkExecutionService}
}

// Execute usa o serviço de execução para processar todos os targets informados.
func (u *runBatchCheckUseCase) Execute(ctx context.Context, cmd command.RunCheckBatchCommand) error {
	return u.checkExecutionService.RunBatch(ctx, cmd.Targets)
}
