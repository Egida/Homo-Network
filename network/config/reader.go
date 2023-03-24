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
	Api struct {
		Server            string
		Port              string
		CustomPathEnabled bool
		CustomPath        string
	}
	Logging struct {
		BotToken string
		ChatId   string
		Logging  bool
	}
	InjectFile struct {
		Linux string
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
