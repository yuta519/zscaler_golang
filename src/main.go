package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"zscaler_golang/auth"
	"zscaler_golang/config"
)

type Payload struct {
	APIKey    string `json:"apiKey"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Timestamp string `json:"timestamp"`
}

func login(
	hostname string,
	api_key string,
	username string,
	password string,
	timestamp int,
) {
	base, _ := url.Parse("https://" + hostname)
	reference, _ := url.Parse("/api/v1/authenticatedSession")
	endpoint := base.ResolveReference(reference).String()

	payload := new(Payload)
	payload.APIKey = api_key
	payload.Username = username
	payload.Password = password
	payload.Timestamp = strconv.Itoa(timestamp)
	payload_json, _ := json.Marshal(payload)

	fmt.Printf("[+] %s\n", string(payload_json))

	res, err := http.Post(endpoint, "application/json", bytes.NewBuffer(payload_json))
	if err != nil {
		fmt.Println("Request Error: ", err)
		return
	}
	defer res.Body.Close()
	fmt.Println(res)
}

func main() {
	api_key := auth.Auth.ObfuscatedApiKey
	timestamp := auth.Auth.Timestamp
	hostname := config.Config.Hostname
	username := config.Config.UserName
	password := config.Config.Password
	login(hostname, api_key, username, password, timestamp)
	// fmt.Println(hostname, username, password, login_session)
}
