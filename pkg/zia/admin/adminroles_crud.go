package admin

import (
	"encoding/json"
	"net/url"

	"zscaler_golang/config"
	"zscaler_golang/pkg/infra"
	"zscaler_golang/pkg/zia/auth"
)

type AdminRole struct {
	Id       int    `json:"id"`
	Rank     int    `json:"rank"`
	Name     string `json:"name"`
	RoleType string `json:"roleType"`
}

func FetchAllAdminRoles() []AdminRole {
	baseUrl, _ := url.Parse("https://" + config.Config.Hostname)
	adminroleRef, _ := url.Parse("/api/v1/adminRoles/lite")
	adminroleEndpoint := baseUrl.ResolveReference(adminroleRef).String()

	session_id := auth.Login()
	response := infra.GetApi(adminroleEndpoint, session_id)
	auth.Logout()

	var adminRoles []AdminRole
	json.Unmarshal(response, &adminRoles)
	return adminRoles
}
