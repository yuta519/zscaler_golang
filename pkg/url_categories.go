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

// type OptAllUrlCategories struct {
// 	IsUrls                             bool
// 	IsDbCategorizedUrls                bool
// 	IsCustomCateogry                   bool
// 	IsEditable                         bool
// 	IsDescription                      bool
// 	IsType                             bool
// 	IsVal                              bool
// 	IsCustomUrlsCount                  bool
// 	IsUrlsRetainingParentCategoryCount bool
// }

// type optionAll func(*OptAllUrlCategories)

// func OptionAllUrlCategories(
// 	is_url bool,
// 	is_db_categorizedUrls bool,
// 	IsCustomCateogry
// 	IsEditable
// 	IsDescription
// 	IsType
// 	IsVal
// 	IsCustomUrlsCount
// 	IsUrlsRetainingParentCategoryCount

// ) {
// 	return func(opt *OoptionAll) {
// 		opt.IsUrls                            =
// 		opt.IsDbCategorizedUrls               =
// 		opt.IsCustomCateogry                  =
// 		opt.IsEditable                        =
// 		opt.IsDescription                     =
// 		opt.IsType                            =
// 		opt.IsVal                             =
// 		opt.IsCustomUrlsCount                 =
// 		opt.IsUrlsRetainingParentCategoryCount=
// 	}
// }

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
		// fmt.Println(url_category.ConfiguredName)
		// if url_category.ConfiguredName != "" {
		// 	shaped_results = append(shaped_results, url_category.Id)
		// } else {
		// 	shaped_results = append(shaped_results, url_category.ConfiguredName)
		// }
	}
	fmt.Println(category_names)
	// fmt.Print(shaped_results)
	return url_categories
}

func FetchLightWeightUrlCategories() []byte {
	session_id := Login()
	is_full := false
	response := GetApi(endpoint_rul_categories(is_full), session_id)
	Logout()
	return response
}
