package providers

import (
	"github.com/brunobotter/site-sentinel/api/controllers"
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

}
