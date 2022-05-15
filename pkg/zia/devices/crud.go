package devices

import (
	"encoding/json"
	"net/url"
	"zscaler_golang/pkg/infra"
	"zscaler_golang/pkg/zia/auth"
	"zscaler_golang/pkg/zia/config"
)

type DeviceGroup struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	GroupType   string `json:"groupType"`
	Description string `json:"description"`
	OsType      string `json:"osType"`
	Predefined  bool   `json:"predefined"`
}

func FetchDeviceGroups() []DeviceGroup {
	sessionId := auth.Login()
	baseUrl, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/deviceGroups")
	endpoint := baseUrl.ResolveReference(reference).String()
	response := infra.GetApi(endpoint, sessionId)
	auth.Logout()

	var deviceGroups []DeviceGroup
	json.Unmarshal(response, &deviceGroups)
	return deviceGroups
}
