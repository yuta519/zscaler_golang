package network

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"zscaler_golang/pkg/zia/auth"
	"zscaler_golang/pkg/zia/config"
)

type IpGroup struct {
	Id             int      `json:"id"`
	Name           string   `json:"name"`
	Type           string   `json:"type"`
	Addresses      []string `json:"addresses"`
	Description    string   `json:"description"`
	CreatorContext string   `json:"creatorContext"`
	Countries      []string `json:"countries"`
	UrlCategories  []string `json:"urlCategories"`
	IpAddresses    []string `json:"ipAddresses"`
}

func FetchIpDstGroups() string {
	baseUrl, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/ipDestinationGroups")
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
	var ipGroups []IpGroup
	json.Unmarshal(byteArray, &ipGroups)
	// fmt.Print(ipGroups)
	return string(byteArray)
}

func FetchIpSrcGroups() string {
	baseUrl, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/ipSourceGroups")
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
	var ipGroups []IpGroup
	json.Unmarshal(byteArray, &ipGroups)
	// fmt.Print(ipGroups)
	return string(byteArray)
}
