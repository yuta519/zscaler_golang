package auth

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gopkg.in/ini.v1"
)

type Auth struct {
	ApiKey string
}

var auth Auth

func init() {
	unix_now := time.Now().UnixNano() / int64(time.Millisecond)
	convert_str_unix := strconv.FormatInt(unix_now, 10)
	key_from_unix := convert_str_unix[len(convert_str_unix)-6:]
	r, _ := strconv.Atoi(key_from_unix)
	shifted_key := fmt.Sprintf("%06d", r>>1)

	apikey := auth.ApiKey
	var obfuscatedApiKey string

	for _, i := range key_from_unix {
		index, _ := strconv.Atoi(string(i))
		obfuscatedApiKey += string(apikey[index])
	}

	for _, i := range shifted_key {
		index, _ := strconv.Atoi(string(i))
		obfuscatedApiKey += string(apikey[index])
	}

	fmt.Println(obfuscatedApiKey)
	return obfuscatedApiKey
	auth = Auth{
		ApiKey: obfuscatedApiKey,
	}

}

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
