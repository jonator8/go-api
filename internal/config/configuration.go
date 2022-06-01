package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App AppConfig
}

type AppConfig struct {
	Host string
	Port string
}

func getAppConfiguration(v *viper.Viper) AppConfig {
	return AppConfig{
		Host: v.GetString("host.address"),
		Port: v.GetString("host.port"),
	}
}

func GetConfiguration() (Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath("./internal/config")

	err := v.ReadInConfig()
	if err != nil {
		return Config{}, fmt.Errorf("error reading in config: %s", err)
	}

	return Config{
		App: getAppConfiguration(v),
	}, nil
}
