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
	session_id := Login()
	url_base, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/urlFilteringRules")
	endpoint := url_base.ResolveReference(reference).String()

	var payload CreateUrlFilterRuleParameter
	payload.AccessControl = accessControl
	payload.Name = name
	payload.Order = order
	payload.Protocols = protocols
	payload.Users = users
	// payload.Uers = make([]interface{},
	// 	"id":   43270458,
	// 	"name": "Aoyama Naoki(naoki.aoyama02@casb89.onmicrosoft.com)",
	// })
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
	// // payload.Uers = make([]interface{},
	// // 	"id":   43270458,
	// // 	"name": "Aoyama Naoki(naoki.aoyama02@casb89.onmicrosoft.com)",
	// // })
	// payload.UrlCategories = []string{"CUSTOM_06", "CUSTOM_09"}
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
	Logout()

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
	// 	"users":[
	// 		{
	// 			"id":43270458,
	// 			"name":"Aoyama Naoki(naoki.aoyama02@casb89.onmicrosoft.com)"
	// 		}
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
