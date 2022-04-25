package firewall

import (
	"encoding/json"
	"net/url"

	"zscaler_golang/pkg/infra"
	"zscaler_golang/pkg/zia/auth"
	"zscaler_golang/pkg/zia/config"
)

type FwRule struct {
	AccessControle    string   `json:"accessControl"`
	EnableFullLogging bool     `json:"enableFullLogging"`
	Id                int      `json:"id"`
	Name              string   `json:"name"`
	Order             int      `json:"order"`
	Rank              int      `json:"rank"`
	Action            string   `json:"action"`
	State             string   `json:"state"`
	DestIpCategories  []string `json:"destIpCategories"`
	ResCategories     []string `json:"resCategories"`
	DestCountries     []string `json:"destCountries"`
	DefaultRule       bool     `json:"defaultRule"`
	Predefined        bool     `json:"predefined"`
}

func FetchAllFwRules() []FwRule {
	sessionId := auth.Login()
	baseUrl, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/firewallFilteringRules")
	endpoint := baseUrl.ResolveReference(reference).String()
	response := infra.GetApi(endpoint, sessionId)
	auth.Logout()

	var fwRules []FwRule
	json.Unmarshal(response, &fwRules)

	return fwRules
}

func FetchSpecificFwRule(id string) FwRule {
	sessionId := auth.Login()
	baseUrl, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/firewallFilteringRules/" + id)
	endpoint := baseUrl.ResolveReference(reference).String()
	response := infra.GetApi(endpoint, sessionId)
	auth.Logout()

	var fwRule FwRule
	json.Unmarshal(response, &fwRule)

	return fwRule
}
