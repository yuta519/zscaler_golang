package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"zscaler_golang/config"
)

type UrlFilteringRule struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Order     int      `json:"order"`
	Protocol  []string `json:"protocol"`
	Locations []string `json:"locations"`
	// TODO: later...
}

func FetchAllUrlFilteringRules() {
	session_id := Login()
	url_base, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/urlFilteringRules")
	endpoint := url_base.ResolveReference(reference).String()
	response := GetApi(endpoint, session_id)
	Logout()
	fmt.Println(string(response))
}

func FetchSpecifiedUrlFilteringRule() UrlFilteringRule {
	// session_id := Login()
	// url_base, _ := url.Parse("https://" + config.Config.Hostname)
	// reference, _ := url.Parse("/api/v1/urlFilteringRules")
	// endpoint := url_base.ResolveReference(reference).String()
	// response := GetApi(endpoint, session_id)
	// Logout()

	var rule UrlFilteringRule
	return rule
}

func CreateUrlFilteringRule(
	id string,

) {
	session_id := Login()
	url_base, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/urlFilteringRules")
	endpoint := url_base.ResolveReference(reference).String()
	var payload CreateAdminUserParameter
	payload_json, _ := json.Marshal(payload)

	req, _ := http.NewRequest(
		"POST",
		endpoint,
		bytes.NewBuffer(payload_json),
	)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("cookie", session_id)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	Logout()

	resp_byte, _ := ioutil.ReadAll(resp.Body)
	var message string
	if resp.StatusCode == 200 {
		var url_filtering_rule SuccessResultOfCreatedAdminUser
		json.Unmarshal(resp_byte, &url_filtering_rule)
		message = "Success: " + url_filtering_rule.LoginName + " is created."
	} else {
		var failed_message map[string]string
		json.Unmarshal(resp_byte, &failed_message)
		message = "Failed: " + strconv.Itoa(resp.StatusCode) + failed_message["code"] + failed_message["message"]
	}
	fmt.Print(string(message))
}

func UpdateFilteringRule() {
}
