package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetApi(url string, session_id string) string {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("cookie", session_id)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	byteArray, _ := ioutil.ReadAll(resp.Body)
	return string(byteArray)
}

func PostApi(url string, session_id string) string {
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("cookie", session_id)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	byteArray, _ := ioutil.ReadAll(resp.Body)
	return string(byteArray)
}
