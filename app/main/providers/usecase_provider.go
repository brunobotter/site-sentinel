package providers

import (
	"github.com/brunobotter/site-sentinel/application/service"
	"github.com/brunobotter/site-sentinel/application/usecase"
	"github.com/brunobotter/site-sentinel/main/container"
)

type UseCaseProvider struct{}

func NewUseCaseProvider() *UseCaseProvider {
	return &UseCaseProvider{}
}
func (p *UseCaseProvider) Register(c container.Container) {
	c.Singleton(func(targetService service.TargetService) usecase.CreateTargetUseCase {
		return usecase.NewCreateTargetUseCase(targetService)
	})

	c.Singleton(func(targetService service.TargetService) usecase.ListTargetsUseCase {
		return usecase.NewListTargetsUseCase(targetService)
	})

	c.Singleton(func(resultService service.CheckResultService) usecase.ListLatestResultsUseCase {
		return usecase.NewListLatestResultsUseCase(resultService)
	})

	c.Singleton(func(checkExecutionService service.CheckExecutionService) usecase.RunBatchCheckUseCase {
		return usecase.NewRunBatchCheckUseCase(checkExecutionService)
	})
}
