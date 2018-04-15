package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

//Config struct, must be injected where needed.
type Config struct {
	secret string
}

//InitConfig reads configuration files
func InitConfig() (*Config, error) {

	var (
		secret string
	)

	if os.Getenv("ENV") == "PRODUCTION" {
		viper.SetEnvPrefix("BONAFIDE")
		viper.AutomaticEnv()
		fmt.Println("Using Env Config")

		secret = viper.Get("secret").(string)
	} else {
		fmt.Println("Using Conf File")
		viper.SetConfigName(".conf")
		viper.SetConfigType("toml")
		viper.AddConfigPath(filepath.Dir(""))

		err := viper.ReadInConfig()
		if err != nil {
			return nil, errors.Wrap(err, "Trying to read config")
		}

		secret = viper.GetString("bonafide.secret")

	}

	config := NewConfig(secret)

	//create ./tmp folder on startup
	if _, err := os.Stat("./tmp"); os.IsNotExist(err) {
		os.Mkdir("./path", os.ModePerm)
	}

	return config, nil
}

//NewConfig Returns a new configuration object
func NewConfig(vals ...string) *Config {
	return &Config{
		secret: vals[0],
	}
}
