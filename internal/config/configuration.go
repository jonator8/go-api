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

func getDataBaseConfiguration(v *viper.Viper) DataBaseConfig {
	return DataBaseConfig{
		Host:     v.GetString("database.host"),
		Port:     v.GetString("database.port"),
		User:     v.GetString("database.user"),
		Password: v.GetString("database.password"),
		Database: v.GetString("database.name"),
	}
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
		Db:  getDataBaseConfiguration(v),
	}, nil
}
