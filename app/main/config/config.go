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

	v.BindEnv("server.port", "SERVER_PORT", "PORT")
	v.BindEnv("server.host", "SERVER_HOST", "HOST")
	v.BindEnv("app_name", "APP_NAME")
	v.BindEnv("env", "env")
	v.BindEnv("database.url", "DATABASE_URL")
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
