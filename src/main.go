package main

import (
	"fmt"
	"zscaler_golang/auth"
	"zscaler_golang/config"
)

func main() {
	login_session := auth.Auth.ObfuscatedApiKey
	hostname := config.Config.Hostname
	username := config.Config.UserName
	password := config.Config.Password
	fmt.Println(hostname, username, password, login_session)
}
