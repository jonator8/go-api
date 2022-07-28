package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App AppConfig
	Db  DataBaseConfig
}

type DataBaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}
type AppConfig struct {
	Host string
	Port string
}

func getDataBaseConfiguration() DataBaseConfig {
	return DataBaseConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.name"),
	}
}

func getAppConfiguration() AppConfig {
	return AppConfig{
		Host: viper.GetString("host.address"),
		Port: viper.GetString("host.port"),
	}
}

func GetConfiguration() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./internal/config")

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, fmt.Errorf("error reading in config: %s", err)
	}

	return Config{
		App: getAppConfiguration(),
		Db:  getDataBaseConfiguration(),
	}, nil
}
