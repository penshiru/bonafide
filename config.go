package main

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

//Config struct, must be injected where needed.
type Config struct {
	Secret string
}

//InitConfig reads configuration files
func InitConfig() (*Config, error) {

	viper.SetEnvPrefix("bdf")
	if os.Getenv("Enviroment") == "Dev" {
		viper.SetConfigName(".conf")
		viper.SetConfigType("toml")
		viper.AddConfigPath(filepath.Dir(""))
		viper.ReadInConfig()
	} else {
		viper.AutomaticEnv()
	}

	//defaults
	viper.SetDefault("BDF_SECRET", "Random string")

	secret := viper.GetString("BDF_SECRET")

	return NewConfig(secret), nil
}

//NewConfig Returns a new configuration object
func NewConfig(vals ...interface{}) *Config {
	return &Config{
		Secret: vals[0].(string),
	}
}
