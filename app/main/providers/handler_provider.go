package providers

import (
	"github.com/brunobotter/site-sentinel/api/controllers"
	"github.com/brunobotter/site-sentinel/application/usecase"
	"github.com/brunobotter/site-sentinel/main/container"
)

type ControllerProvider struct{}

func NewControllereProvider() *ControllerProvider {
	return &ControllerProvider{}
}
func (p *ControllerProvider) Register(c container.Container) {
	c.Singleton(func() *controllers.HealthHandler {
		return controllers.NewHealthHandler()
	})
	c.Singleton(func(
		createTargetUseCase usecase.CreateTargetUseCase,
		listTargetsUseCase usecase.ListTargetsUseCase,
		runBatchCheckUseCase usecase.RunBatchCheckUseCase,
		listLatestResultUseCase usecase.ListLatestResultsUseCase,
	) *controllers.MonitorHandler {
		return controllers.NewMonitorHandler(
			createTargetUseCase,
			listTargetsUseCase,
			runBatchCheckUseCase,
			listLatestResultUseCase,
		)
	})
}
