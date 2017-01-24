package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

func LoadFromFile(filePath string) *Config {
	var c Config
	if _, err := toml.DecodeFile(filePath, &c); err != nil {
		log.Fatalf("Failed to read config file: %s", err.Error())
	}
	return &c
}
