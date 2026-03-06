package providers

import (
	"github.com/brunobotter/site-sentinel/infra/logger"
	"github.com/brunobotter/site-sentinel/main/config"
	"github.com/brunobotter/site-sentinel/main/container"
)

type ConfigServiceProvider struct{}

func NewConfigServiceProvider() *ConfigServiceProvider {
	return &ConfigServiceProvider{}
}

func (p *ConfigServiceProvider) Register(c container.Container) {
	c.Singleton(func() *config.Config {
		cfg := config.Init()
		return cfg
	})

	c.Singleton(func(cfg *config.Config) logger.Logger {
		return logger.NewJammesLogger(cfg.App_Name, cfg.Env, false)
	})
}
