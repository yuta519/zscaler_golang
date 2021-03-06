package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"zscaler_golang/pkg/zia/auth"
	"zscaler_golang/pkg/zia/config"
)

func FetchAllNetworkServices() string {
	baseUrl, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/networkServices")
	endpoint := baseUrl.ResolveReference(reference).String()

	sessionId := auth.Login()
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("cookie", sessionId)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	auth.Logout()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	return string(byteArray)
}
