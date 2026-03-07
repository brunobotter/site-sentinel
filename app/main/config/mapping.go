package config

type Config struct {
	Server   ServerConfig   `mapstruct:"server"`
	App_Name string         `mapstruct:"app_name"`
	Env      string         `mapstruct:"env"`
	Database DatabaseConfig `mapstruct:"database"`
	Monitor  MonitorConfig  `mapstruct:"monitor"`
}

type ServerConfig struct {
	Port int    `mapstruct:"port"`
	Host string `mapstruct:"host"`
}

type DatabaseConfig struct {
	URL string `mapstruct:"url"`
}

type MonitorConfig struct {
	Enabled         bool `mapstruct:"enabled"`
	Workers         int  `mapstruct:"workers"`
	JobQueue        int  `mapstruct:"job_queue"`
	IntervalSeconds int  `mapstruct:"interval_seconds"`
}
