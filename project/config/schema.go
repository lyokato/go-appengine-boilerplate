package config

import "fmt"

type Config struct {
	Web *WebConfig `toml:"web"`
}

type WebConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

func (wc *WebConfig) Address() string {
	return fmt.Sprintf("%s:%d", wc.Host, wc.Port)
}
