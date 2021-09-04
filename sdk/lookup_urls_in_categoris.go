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

func LookupUrlCategory(tareget_urls []string) string {
	session_id := Login()

	url_base, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/urlLookup")
	endpoint := url_base.ResolveReference(reference).String()

	payload := tareget_urls
	payload_json, _ := json.Marshal(payload)

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
}
