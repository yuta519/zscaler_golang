package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetApi(url string, session_id string) []byte {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("cookie", session_id)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func PostApi(url string, session_id string) []byte {
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("cookie", session_id)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
