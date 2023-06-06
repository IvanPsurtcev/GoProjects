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

	Db struct {
		Host            string `json:"host"`
		Port            string `json:"port"`
		User            string `json:"user"`
		Password        string `json:"password"`
		DBName          string `json:"dbName"`
		SSLMode         string `json:"sslMode"`
		MaxOpenConns    int    `json:"maxOpenConns"`
		MaxIdleConns    int    `json:"maxIdleConns"`
		ConnMaxIdleTime int    `json:"connMaxIdleTimeSec"`
	} `json:"db"`

	ReservoirApiUrl string `json:"reservoirApiUrl"`

	Jwt struct {
		Duration   int64  `json:"duration"`
		Issuer     string `json:"issuer"`
		HmacSecret string `json:"hmacSecret"`
	} `json:"jwt"`

	Contract struct {
		OwnerAddress     string `json:"ownerAddress"`
		OwnerPK          string `json:"ownerPK"`
		ModeratorAddress string `json:"moderatorAddress"`
		ModeratorPK      string `json:"moderatorPK"`
		ContractAddress  string `json:"contractAddress"`
		ChainID          int64  `json:"chainID"`
		Rpc              string `json:"rpc"`
	} `json:"contract"`
}

func (c Config) GetDsn() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Db.Host, c.Db.Port,
		c.Db.User, c.Db.Password,
		c.Db.DBName, c.Db.SSLMode)
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
