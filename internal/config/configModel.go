package config

import "github.com/ruziba3vich/configurations/internal/models"

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Albums []models.Album `yaml:"albums"`
}

func New(host, port string) *Config {
	return &Config{
		Server: struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
		}{
			Host: "localhost",
			Port: "7777",
		},
		Albums: []models.Album{},
	}
}
