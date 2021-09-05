package pkg

import (
	"fmt"
	"net/url"
	"zscaler_golang/config"
)

type UrlFilteringRule struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Order     int      `json:"order"`
	Protocol  []string `json:"protocol"`
	Locations []string `json:"locations"`
	// TODO: later...
}

func FetchAllUrlFilteringRules() {
	session_id := Login()

	url_base, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/urlFilteringRules")
	endpoint := url_base.ResolveReference(reference).String()

	resposne := GetApi(endpoint, session_id)

	fmt.Println(string(resposne))
}

func FetchSpecifiedUrlFilteringRule() UrlFilteringRule {
	var rule UrlFilteringRule
	return rule
}

func CreateUrlFilteringRule() {
}

func UpdateFilteringRule() {
}
