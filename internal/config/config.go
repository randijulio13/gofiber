package config

import "github.com/spf13/viper"

type ConfigStruct struct {
	AppConfig      `mapstructure:",squash"`
	DatabaseConfig `mapstructure:",squash"`
	JwtConfig      `mapstructure:",squash"`
}

var config ConfigStruct

func GetConfig() *ConfigStruct {
	return &config
}

func InitializeConfig() *ConfigStruct {
	viper.Unmarshal(&config)

	return &config
}

func SetDefaultConfig() {
	viper.SetDefault("APP_NAME", "API ECOMMERCE")
	viper.SetDefault("APP_PORT", 3000)

	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", 3306)
	viper.SetDefault("DB_USER", "root")
	viper.SetDefault("DB_PASSWORD", "")
	viper.SetDefault("DB_NAME", "api_ecomm")

	viper.SetDefault("JWT_ACCESS_SECRET", "")
	viper.SetDefault("JWT_REFRESH_SECRET", "")
}
