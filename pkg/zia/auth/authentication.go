package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"zscaler_golang/pkg/zia/config"
)

type ApiCredential struct {
	APIKey    string `json:"apiKey"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Timestamp string `json:"timestamp"`
}

type AuthPrepare struct {
	ObfuscatedApiKey string
	Timestamp        int
}

var auth AuthPrepare

func init() {
	unix_now := time.Now().UnixNano() / int64(time.Millisecond)
	timestamp := int(unix_now)
	convert_str_unix := strconv.FormatInt(unix_now, 10)
	key_from_unix := convert_str_unix[len(convert_str_unix)-6:]
	r, _ := strconv.Atoi(key_from_unix)
	shifted_key := fmt.Sprintf("%06d", r>>1)

	apikey := config.Config.ApiKey
	var obfuscatedApiKey string
	for _, i := range key_from_unix {
		index, _ := strconv.Atoi(string(i))
		obfuscatedApiKey += string(apikey[index])
	}
	for _, i := range shifted_key {
		index, _ := strconv.Atoi(string(i))
		obfuscatedApiKey += string(apikey[index+2])
	}

	auth = AuthPrepare{
		ObfuscatedApiKey: obfuscatedApiKey,
		Timestamp:        timestamp,
	}
}

func Login() string {
	var session_id string

	base, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/authenticatedSession")
	endpoint := base.ResolveReference(reference).String()
	payload := new(ApiCredential)
	payload.APIKey = auth.ObfuscatedApiKey
	payload.Username = config.Config.UserName
	payload.Password = config.Config.Password
	payload.Timestamp = strconv.Itoa(auth.Timestamp)
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
