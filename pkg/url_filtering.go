package pkg

import (
	"encoding/json"
	"fmt"
)

type UrlFilteringRule struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Order     int      `json:"order"`
	Protocol  []string `json:"protocol"`
	Locations []string `json:"locations"`
	// TODO: later...
}

func FetchAllUrlFilteringRules() []UrlFilteringRule {
	session_id := Login()
	is_full := true
	response := GetApi(endpoint_rul_categories(is_full), session_id)
	Logout()
	var url_categories []UrlCategory
	json.Unmarshal(response, &url_categories)
	for i := range url_categories {
		fmt.Println(url_categories[i].Id)
	}
	var rules []UrlFilteringRule
	return rules
}

func FetchSpecifiedUrlFilteringRule() UrlFilteringRule {
	var rule UrlFilteringRule
	return rule
}

func CreateUrlFilteringRule() {
}

func UpdateFilteringRule() {
}
