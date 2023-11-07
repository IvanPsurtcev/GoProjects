package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
		Cors struct {
			Origin       string `json:"origin"`
			AllowMethods string `json:"allowMethods"`
			AllowHeaders string `json:"allowHeaders"`
		} `json:"cors"`
		Upload struct {
			UploadDir string   `json:"uploadDir"`
			MaxSize   int64    `json:"maxSize"`
			Formats   []string `json:"formats"`
		}
	} `json:"server"`

	Contract struct {
		CollectionFactoryAddress string `json:"collectionFactoryAddress"`
		CollectionAddress        string `json:"collectionAddress"` // TODO: hardcode
		ChainID                  int64  `json:"chainID"`
		Rpc                      string `json:"rpc"`
	} `json:"contract"`
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
