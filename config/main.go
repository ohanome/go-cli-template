package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"go-cli-template/setup"
	"os"
)

type Config struct {
	Sample string `json:"sample"`
}

var defaultConfig = Config{
	Sample: "sample",
}

func init() {
	viper.SetConfigFile(setup.GetConfigFile())
}

func Get(key string) interface{} {
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(err)
	}
	return viper.Get(key)
}

func Set(key string, value interface{}) {
	viper.Set(key, value)
	err := viper.WriteConfig()
	if err != nil {
		panic(err)
	}
}

func SafeDefaults(warnOnExist bool, forceRewrite bool) {
	configFile := setup.GetConfigFile()
	_, err := os.Stat(configFile)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	} else {
		if forceRewrite {
			fmt.Printf("Config file %s will be overwritten.\n", setup.GetConfigFile())
			err = os.Remove(configFile)
			if err != nil {
				panic(err)
			}
		} else {
			if warnOnExist {
				fmt.Printf("Config file already exists and will not be overwritten. If you want to refresh it, delete the old file at %s.\n", setup.GetConfigFile())
			}
			return
		}
	}

	configJson, err := json.Marshal(defaultConfig)
	if err != nil {
		panic(err)
	}

	err = viper.ReadConfig(bytes.NewBuffer(configJson))
	if err != nil {
		panic(err)
	}

	viper.SetConfigFile(setup.GetConfigFile())
	err = viper.WriteConfig()
	if err != nil {
		panic(err)
	}
}
