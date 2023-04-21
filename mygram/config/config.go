package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	var cfg Config
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s\n", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("Error unmarshalling config: %s\n", err)
	}

	return &cfg, nil
}
