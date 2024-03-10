package apiserver

import (
	"encoding/json"
	"log"
	"os"

	. "github.com/ElnurKoke/web-api.git/internal/model"
)

func NewConfig() (*Config, error) {
	configFile, err := os.Open("config/apiserver.json")
	if err != nil {
		return &Config{}, err
	}
	defer configFile.Close()
	var config Config
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return &Config{}, err
	}
	log.Println("Configuration extraction successful")
	return &config, nil
}
