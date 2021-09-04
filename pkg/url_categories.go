package pkg

import (
	"encoding/json"
	"fmt"
	"net/url"
	"zscaler_golang/config"
)

func endpoint_rul_categories(is_full bool) string {
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

type UrlCategory struct {
	Id                               string   `json:"id"`
	Urls                             []string `json:"urls"`
	DbCategorizedUrls                []string `json:"dbCategorizedUrls"`
	CustomCateogry                   bool     `json:"customCateogry"`
	Editable                         bool     `json:"editable"`
	Description                      string   `json:"description"`
	Type                             string   `json:"Type"`
	Val                              int      `json:"val"`
	CustomUrlsCount                  int      `json:"customUrlsCount"`
	UrlsRetainingParentCategoryCount int      `json:"urlsRetainingParentCategoryCount"`
}

func FetchAllUrlCategories() []UrlCategory {
	session_id := Login()
	is_full := true
	response := GetApi(endpoint_rul_categories(is_full), session_id)
	Logout()
	var url_categories []UrlCategory
	json.Unmarshal(response, &url_categories)
	for i := range url_categories {
		fmt.Println(url_categories[i].Id)
	}
	return url_categories
}

func FetchLightWeightUrlCategories() []byte {
	session_id := Login()
	is_full := false
	response := GetApi(endpoint_rul_categories(is_full), session_id)
	Logout()
	return response
}

// func LookupUrlCategory(tareget_urls []string) string {
// 	session_id := Login()

// 	url_base, _ := url.Parse("https://" + config.Config.Hostname)
// 	reference, _ := url.Parse("/api/v1/urlLookup")
// 	endpoint := url_base.ResolveReference(reference).String()

// 	payload := tareget_urls
// 	payload_json, _ := json.Marshal(payload)

// 	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(payload_json))
// 	req.Header.Set("content-type", "application/json")
// 	req.Header.Set("cache-control", "no-cache")
// 	req.Header.Set("cookie", session_id)

// 	client := new(http.Client)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	Logout()
// 	byteArray, _ := ioutil.ReadAll(resp.Body)
// 	return string(byteArray)
// }
