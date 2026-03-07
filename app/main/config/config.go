package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func Init() *Config {
	cfg, err := Read()
	if err != nil {
		panic(fmt.Sprintf("Erro ao ler configuração de arquivo: %v", err))
	}
	return cfg
}

func Read() (*Config, error) {
	v := viper.New()
	v.SetConfigFile(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetDefault("server.port", 8080)
	v.SetDefault("server.host", "0.0.0.0")

	v.SetDefault("monitor.enabled", true)
	v.SetDefault("monitor.workers", 300)
	v.SetDefault("monitor.job_queue", 20000)
	v.SetDefault("monitor.interval_seconds", 60)

	v.BindEnv("server.port", "SERVER_PORT", "PORT")
	v.BindEnv("server.host", "SERVER_HOST", "HOST")
	v.BindEnv("app_name", "APP_NAME")
	v.BindEnv("env", "env")
	v.BindEnv("database.url", "DATABASE_URL")
	v.BindEnv("monitor.interval_seconds", "MONITOR_INTERVAL_SECONDS")
	v.BindEnv("monitor.job_queue", "MONITOR_JOB_QUEUE")
	v.BindEnv("monitor.workers", "MONITOR_WORKERS")
	v.BindEnv("monitor.enabled", "MONITOR_ENABLED")
	err := v.ReadInConfig()
	if err != nil {
		var configNotFound viper.ConfigFileNotFoundError
		if !errors.As(err, &configNotFound) && !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
	}

	conf := Config{}
	err = v.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil

}
