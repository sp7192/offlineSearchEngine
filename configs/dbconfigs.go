package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type DbConfigs struct {
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	IP       string `mapstructure:"DB_IP"`
	Port     uint   `mapstructure:"DB_PORT"`
	Name     string `mapstructure:"DB_NAME"`
}

func LoadDbConfigs(path string) (Configs, error) {
	var ret Configs
	viper.AddConfigPath(path)
	viper.SetConfigName("dbconfigs")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	err = viper.Unmarshal(&ret)
	return ret, err
}
