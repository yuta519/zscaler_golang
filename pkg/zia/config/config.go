package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	UserName string
	Password string
	Hostname string
	ApiKey   string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		UserName: cfg.Section("credential").Key("USERNAME").String(),
		Password: cfg.Section("credential").Key("PASSWORD").String(),
		Hostname: cfg.Section("credential").Key("HOSTNAME").String(),
		ApiKey:   cfg.Section("credential").Key("APIKEY").String(),
	}
}
