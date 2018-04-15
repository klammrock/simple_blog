package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"simple_blog/blog_service"
	"simple_blog/db"
)

var (
	isDev = true
)

type Config struct {
	// "DEV" or "PROD"
	Config      string                   `json:"config"`
	ServicePort int                      `json:"service_port"`
	Db          db.Config                `json:"db"`
	Admin       blog_service.AdminConfig `json:"admin"`
}

func IsDev() bool {
	return isDev
}

func LoadConfig(configFilePath string) (*Config, error) {
	f, err := os.Open(configFilePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	err = setIsDev(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func setIsDev(c *Config) error {
	if c.Config == "DEV" {
		isDev = true
	} else if c.Config == "PROD" {
		isDev = false
	} else {
		return fmt.Errorf("config should be DEV or PROD")
	}

	return nil
}
