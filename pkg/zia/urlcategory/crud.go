package urlcategory

import (
	"encoding/json"
	"fmt"
	"net/url"

	"zscaler_golang/pkg/infra"
	"zscaler_golang/pkg/zia/auth"
	"zscaler_golang/pkg/zia/config"
)

type UrlCategory struct {
	Id                               string   `json:"id"`
	ConfiguredName                   string   `json:"configuredName"`
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
	session_id := auth.Login()
	is_full := true
	response := infra.GetApi(endpoint_rul_categories(is_full), session_id)
	auth.Logout()

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

	session_id := auth.Login()
	is_full := false
	response := infra.GetApi(endpoint_rul_categories(is_full), session_id)
	auth.Logout()
	return response
}
