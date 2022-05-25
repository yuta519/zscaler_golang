package config

import (
	"os"
)

type ConfigList struct {
	UserName string
	Password string
	Hostname string
	ApiKey   string
}

var Config ConfigList

func init() {
	Config = ConfigList{
		UserName: os.Getenv("ZIA_USERNAME"),
		Password: os.Getenv("ZIA_PASSWORD"),
		Hostname: os.Getenv("ZIA_HOSTNAME"),
		ApiKey:   os.Getenv("ZIA_APIKEY"),
	}
}
