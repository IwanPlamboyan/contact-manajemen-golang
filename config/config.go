package config

import (
	"github.com/IwanPlamboyan/contact-manajemen-golang/app"
	"github.com/IwanPlamboyan/contact-manajemen-golang/utils"
	"github.com/spf13/viper"
)

type AppConfig struct {
	AppName string `mapstructure:"APP_NAME"`
	AppENV  string `mapstructure:"APP_ENV"`
	JWTSecret string `mapstructure:"JWT_SECRET"`

	DatabaseConfig app.DatabaseConfig `mapstructure:",squash"`
}

func LoadConfig() (*AppConfig, error) {
	config := viper.New()
	config.SetConfigFile("config-local.env")
	config.SetConfigType("env")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	// read config
	if err := config.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg AppConfig
	if err := config.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func ProvideDatabaseConfig(cfg *AppConfig) *app.DatabaseConfig {
	return &cfg.DatabaseConfig
}

func ProvideJWTUtil(cfg *AppConfig) *utils.JWTUtil {
    return utils.NewJWTUtil(cfg.JWTSecret)
}