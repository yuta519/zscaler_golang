package pkg

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

func Login() string {
	var session_id string

	base, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/authenticatedSession")
	endpoint := base.ResolveReference(reference).String()
	payload := new(Payload)
	payload.APIKey = auth.Auth.ObfuscatedApiKey
	payload.Username = config.Config.UserName
	payload.Password = config.Config.Password
	payload.Timestamp = strconv.Itoa(auth.Auth.Timestamp)
	payload_json, _ := json.Marshal(payload)

	res, err := http.Post(endpoint, "application/json", bytes.NewBuffer(payload_json))
	if err != nil {
		fmt.Println("Request Error: ", err)
	}
	defer res.Body.Close()

	for _, cookie := range res.Cookies() {
		if cookie.Name == "JSESSIONID" {
			session_id = "JSESSIONID=" + cookie.Value
		}
	}
	return session_id
}

func Logout() {
	base, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/authenticatedSession")
	endpoint := base.ResolveReference(reference).String()
	_, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		fmt.Println("Request Error: ", err)
	}
}
