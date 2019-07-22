package config

import (
	"github.com/spf13/viper"
)

//Configuration : struct to map from yml or json
type Configuration struct {
	Database Database
	Logger   Logger
	Server   Server
	Redis    Redis
}

//New : initialized configuration
func New() (*Configuration, error) {
	path := "./"

	viper.SetConfigType("yaml")
	viper.SetConfigName("default")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := new(Configuration)
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
