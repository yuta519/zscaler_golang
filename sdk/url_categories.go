package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"zscaler_golang/config"
)

func endpoint(is_full bool) string {
	url_base, _ := url.Parse("https://" + config.Config.Hostname)
	var reference *url.URL
	if is_full {
		reference, _ = url.Parse("/api/v1/urlCategories")
	} else {
		reference, _ = url.Parse("/api/v1/urlCategories/lite")
	}
	endpoint := url_base.ResolveReference(reference).String()
	return endpoint
}

func FetchAllUrlCategories() []byte {
	session_id := Login()
	is_full := true
	response := GetApi(endpoint(is_full), session_id)
	Logout()
	return response
}

func FetchLightWeightUrlCategories() []byte {
	session_id := Login()
	is_full := false
	response := GetApi(endpoint(is_full), session_id)
	Logout()
	return response
}

type UrlLookupPayload struct {
	Urls []string `json:"urls"`
}

func LookupUrlCategory(tareget_urls []string) string {
	session_id := Login()

	url_base, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/urlLookup")
	endpoint := url_base.ResolveReference(reference).String()

	payload := new(UrlLookupPayload)
	payload.Urls = tareget_urls
	// payload := `{"token":"` + tareget_urls `"}`
	payload_json, _ := json.Marshal(payload)
	// fmt.Print(string(payload_json))
	// fmt.Println(tareget_urls)

	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(payload_json))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("cookie", session_id)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	Logout()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	return string(byteArray)
	// return ""
}
