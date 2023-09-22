package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	System struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"system"`
}

func CreateConfig() *Config {
	v := viper.New()

	v.AddConfigPath("./config")
	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var cfg Config
	err = v.Unmarshal(&cfg)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return &cfg
}
