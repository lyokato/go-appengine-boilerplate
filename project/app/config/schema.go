package config

import "fmt"

type Config struct {
	Web          *WebConfig     `toml:"web"`
	AdminSession *SessionConfig `toml:"admin_session"`
	SiteSession  *SessionConfig `toml:"site_session"`
}

type WebConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

func (wc *WebConfig) Address() string {
	return fmt.Sprintf("%s:%d", wc.Host, wc.Port)
}

type SessionConfig struct {
	Name   string `toml:"name"`
	Secret string `toml:"secret"`
	Path   string `toml:"path"`
	Domain string `toml:"domain"`
	MaxAge int    `toml:"max_age"`
	Secure bool   `toml:"secure"`
}
