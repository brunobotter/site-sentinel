package providers

import (
	"github.com/brunobotter/site-sentinel/application/service"
	infraService "github.com/brunobotter/site-sentinel/infra/service"
	"github.com/brunobotter/site-sentinel/main/container"
)

type ServiceProvider struct{}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
func (p *ServiceProvider) Register(c container.Container) {
	c.Singleton(func() service.MonitorPlannerService {
		return infraService.NewMonitorPlannerService(10)
	})

	c.Singleton(func() service.CheckAggregationService {
		return infraService.NewCheckAggregationService()
	})
}
