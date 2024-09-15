package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var Config *viper.Viper

func init() {
	Config = viper.New()
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(fmt.Errorf("error to get the os wd %s", err))
	}
	Config.AddConfigPath(wd + "/config")
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.AutomaticEnv()
	Config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := Config.ReadInConfig(); err != nil {
		fmt.Println(fmt.Errorf("error read in config %s", err))
	}
}
