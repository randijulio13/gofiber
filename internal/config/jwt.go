package config

type JwtConfig struct {
	AccessSecret  string `mapstructure:"jwt_access_secret"`
	RefreshSecret string `mapstructure:"jwt_refresh_secret"`
}
