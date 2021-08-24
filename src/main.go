package main

import (
	"fmt"
	"net/http"
	"zscaler_golang/auth"
	"zscaler_golang/config"
)

func login(hostname string, login_session string) {
	api_endpoint := "https://" + hostname + "/api/v1/status"
	jsession := "JSESSIONID=" + login_session
	req, _ := http.NewRequest("GET", api_endpoint, nil)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("cookie", jsession)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(api_endpoint, resp)
}

func main() {
	login_session := auth.Auth.ObfuscatedApiKey
	hostname := config.Config.Hostname
	// username := config.Config.UserName
	// password := config.Config.Password
	login(hostname, login_session)
	// fmt.Println(hostname, username, password, login_session)
}
