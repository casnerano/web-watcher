package config

import (
	"flag"
	"os"

	"github.com/go-yaml/yaml"
)

var path = "./configs/watcher.yml"

type SMTP struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	SMTP SMTP `yaml:"smtp"`
}

func Load() (*Config, error) {
	c := &Config{}
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func init() {
	flag.StringVar(&path, "config", path, "Config path")
	flag.Parse()
}
