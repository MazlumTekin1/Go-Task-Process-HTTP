package infrastructure

import (
	"encoding/json"
	"os"
)

type Config struct {
	Database DatabaseConfig `json:"database"`
}

type DatabaseConfig struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	User        string `json:"user"`
	Password    string `json:"password"`
	DBName      string `json:"dbname"`
	SSLMode     string `json:"sslmode"`
	MaxPoolSize int    `json:"maxpoolsize"`
}

func LoadConfig(filepath string) (Config, error) {
	var config Config

	file, err := os.Open(filepath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
