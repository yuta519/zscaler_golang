package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"zscaler_golang/config"
)

func FetchAllNetworkServices() string {
	url_base, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/networkServices")
	endpoint := url_base.ResolveReference(reference).String()

	session_id := Login()
	req, _ := http.NewRequest("GET", endpoint, nil)
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
