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

func FetchAllUrlCategories() []UrlCategory {
	session_id := Login()
	is_full := true
	response := GetApi(endpoint_rul_categories(is_full), session_id)
	Logout()

	var url_categories []UrlCategory
	json.Unmarshal(response, &url_categories)

	var category_names []string
	for _, url_category := range url_categories {
		if len(url_category.ConfiguredName) > 0 {
			category_names = append(category_names, url_category.ConfiguredName)
		} else {
			category_names = append(category_names, url_category.Id)
		}
	}
	fmt.Println(category_names)
	return url_categories
}

func FetchLightWeightUrlCategories() []byte {

	session_id := Login()
	is_full := false
	response := GetApi(endpoint_rul_categories(is_full), session_id)
	Logout()
	return response
}
