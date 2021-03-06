package main

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	Port             string `mapstructure:"applicationPort"`
	ConnectionString string `mapstructure:"connetionString"`
}

var AppConfig *config

func LoadAppConfig() {
	log.Println("Loading Server Configuration...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
