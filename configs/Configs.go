package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Configs struct {
	EngineType string `mapstructure:"ENGINE_TYPE"`
	XApiKey    string `mapstructure:"X_API_KEY"`
}

func LoadConfigs(path string) (Configs, error) {
	var ret Configs
	viper.AddConfigPath(path)
	viper.SetConfigName("configs")
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
