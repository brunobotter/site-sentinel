package providers

import (
	apphttp "github.com/brunobotter/site-sentinel/application/http"
	"github.com/brunobotter/site-sentinel/application/repo"
	appservice "github.com/brunobotter/site-sentinel/application/service"
	infraHTTP "github.com/brunobotter/site-sentinel/infra/http"
	"github.com/brunobotter/site-sentinel/infra/logger"
	infraService "github.com/brunobotter/site-sentinel/infra/service"
	"github.com/brunobotter/site-sentinel/main/config"
	"github.com/brunobotter/site-sentinel/main/container"
)

type ServiceProvider struct{}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
func (p *ServiceProvider) Register(c container.Container) {
	c.Singleton(func() appservice.MonitorPlannerService {
		return infraService.NewMonitorPlannerService(10)
	})

	c.Singleton(func() appservice.CheckAggregationService {
		return infraService.NewCheckAggregationService()
	})

	c.Singleton(func(cfg *config.Config, log logger.Logger) apphttp.Client {
		return infraHTTP.NewHttpClient(cfg, log)
	})

	c.Singleton(func(targetRepo repo.MonitorTargetRepository) appservice.TargetService {
		return infraService.NewTargetService(targetRepo)
	})

	c.Singleton(func(resultRepo repo.CheckResultRepository) appservice.CheckResultService {
		return infraService.NewCheckResultService(resultRepo)
	})

	c.Singleton(func(
		planner appservice.MonitorPlannerService,
		httpClient apphttp.Client,
		resultRepo repo.CheckResultRepository,
	) appservice.CheckExecutionService {
		return infraService.NewCheckExecutionService(planner, httpClient, resultRepo)
	})
}
