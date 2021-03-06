package urlfiltering

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"zscaler_golang/pkg/infra"
	"zscaler_golang/pkg/zia/auth"
	"zscaler_golang/pkg/zia/config"
)

type CreateUrlFilterRuleParameter struct {
	AccessControl       string        `json:"accessControl"`
	Name                string        `json:"name"`
	Order               int           `json:"order"`
	Protocols           []string      `json:"protocols"`
	Users               []interface{} `json:"users"`
	UrlCategories       []string      `json:"urlCategories"`
	State               string        `json:"state"`
	Rank                int           `json:"rank"`
	RequestMethods      []string      `json:"requestMethods"`
	BlockOverride       bool          `json:"blockOverride"`
	EnforceTimeValidity bool          `json:"enforceTimeValidity"`
	CbiProfileId        int           `json:"cbiProfileId"`
	Action              string        `json:"action"`
}

type UrlFilteringRule struct {
	Id                  int      `json:"id"`
	AccessControl       string   `json:"accessControl"`
	Name                string   `json:"name"`
	Order               int      `json:"order"`
	Protocols           []string `json:"protocols"`
	UrlCategories       []string `json:"urlCategories"`
	State               string   `json:"state"`
	Rank                int      `json:"rank"`
	RequestMethods      []string `json:"requestMethods"`
	BlockOverride       string   `json:"blockOverride"`
	Description         string   `json:"description"`
	EnforceTimeValidity string   `json:"enforceTimeValidity"`
	CbiProfileId        int      `json:"cbiProfileId"`
	Action              string   `json:"action"`
}

func FetchAllUrlFilteringRules() []UrlFilteringRule {
	sessionId := auth.Login()
	baseUrl, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/urlFilteringRules")
	endpoint := baseUrl.ResolveReference(reference).String()
	response := infra.GetApi(endpoint, sessionId)
	auth.Logout()

	var rules []UrlFilteringRule
	json.Unmarshal(response, &rules)
	return rules
}

func FetchSpecifiedUrlFilteringRule(id string) UrlFilteringRule {
	sessionId := auth.Login()
	baseUrl, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/urlFilteringRules/" + id)
	endpoint := baseUrl.ResolveReference(reference).String()
	response := infra.GetApi(endpoint, sessionId)
	auth.Logout()

	var rule UrlFilteringRule
	json.Unmarshal(response, &rule)
	return rule
}

func CreateUrlFilteringRule(
	accessControl string,
	name string,
	order int,
	protocols []string,
	users []interface{},
	url_categories []string,
	state string,
	rank int,
	request_methods []string,
	block_override bool,
	enforce_time_validity bool,
	cbi_profile_id int,
	action string,
) {
	session_id := auth.Login()
	url_base, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/urlFilteringRules")
	endpoint := url_base.ResolveReference(reference).String()

	var payload CreateUrlFilterRuleParameter
	payload.AccessControl = accessControl
	payload.Name = name
	payload.Order = order
	payload.Protocols = protocols
	payload.Users = users
	payload.UrlCategories = url_categories
	payload.State = state
	payload.Rank = rank
	payload.RequestMethods = request_methods
	payload.BlockOverride = block_override
	payload.EnforceTimeValidity = enforce_time_validity
	payload.CbiProfileId = cbi_profile_id
	payload.Action = action
	payload_json, _ := json.Marshal(payload)

	// payload.AccessControl = "READ_WRITE"
	// payload.Name = "Test Kawamura"
	// payload.Order = 1
	// payload.Protocols = []string{
	// 	"DOHTTPS_RULE",
	// 	"TUNNELSSL_RULE",
	// 	"HTTP_PROXY",
	// 	"FOHTTP_RULE",
	// 	"FTP_RULE",
	// 	"HTTPS_RULE",
	// 	"HTTP_RULE",
	// 	"SSL_RULE",
	// 	"TUNNEL_RULE",
	// }
	// payload.Users = nil
	// payload.State = "ENABLED"
	// payload.Rank = 0
	// payload.RequestMethods = []string{
	// 	"OPTIONS",
	// 	"GET",
	// 	"HEAD",
	// 	"POST",
	// 	"PUT",
	// 	"DELETE",
	// 	"TRACE",
	// 	"CONNECT",
	// 	"OTHER",
	// }
	// payload.BlockOverride = false
	// payload.EnforceTimeValidity = false
	// payload.CbiProfileId = 0
	// payload.Action = "ALLOW"
	// payload_json, _ := json.Marshal(payload)

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
	auth.Logout()

	// byte_response, _ := ioutil.ReadAll(response.Body)

	// return string(byteArray)
	// {
	// 	"id":682031,
	// 	"accessControl":"READ_WRITE",
	// 	"name":"Yahoo Block",
	// 	"order":4,
	// 	"protocols":[
	// 		"DOHTTPS_RULE",
	// 		"TUNNELSSL_RULE",
	// 		"HTTP_PROXY",
	// 		"FOHTTP_RULE",
	// 		"FTP_RULE",
	// 		"HTTPS_RULE",
	// 		"HTTP_RULE",
	// 		"SSL_RULE",
	// 		"TUNNEL_RULE"
	// 	],
	// 	],
	// 	"urlCategories":[
	// 		"CUSTOM_06",
	// 		"CUSTOM_09"
	// 	],
	// 	"state":"ENABLED",
	// 	"rank":7,
	// 	"requestMethods":[
	// 		"OPTIONS","GET","HEAD","POST","PUT","DELETE","TRACE","CONNECT","OTHER"
	// 	],
	// 	"blockOverride":false,
	// 	"enforceTimeValidity":false,
	// 	"cbiProfileId":0,
	// 	"action":"ALLOW"
	// }

	resp_byte, _ := ioutil.ReadAll(resp.Body)
	var message string
	if resp.StatusCode == 200 {
		var url_filtering_rule CreateUrlFilterRuleParameter
		json.Unmarshal(resp_byte, &url_filtering_rule)
		message = "Success: " + url_filtering_rule.Name + " is created."
	} else {
		var failed_message map[string]string
		json.Unmarshal(resp_byte, &failed_message)
		message = "Failed: " + strconv.Itoa(resp.StatusCode) + failed_message["code"] + failed_message["message"]
	}
	fmt.Print(string(message))
}

func UpdateFilteringRule() {
}
