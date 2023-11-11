package config

type AppConfig struct {
	AppName string `mapstructure:"app_name"`
	AppPort string `mapstructure:"app_port"`
}
