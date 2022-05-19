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

type Device struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	DeviceGroupType string `json:"deviceGroupType"`
	DeviceModel     string `json:"deviceModel"`
	OsType          string `json:"osType"`
	OsVersion       string `json:"osVersion"`
	OwnerUserrId    int    `json:"ownerUserId"`
	OwnerName       string `json:"ownerName"`
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

func FetchDevices() []Device {
	sessionId := auth.Login()
	baseUrl, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/deviceGroups/devices")
	endpoint := baseUrl.ResolveReference(reference).String()
	response := infra.GetApi(endpoint, sessionId)
	auth.Logout()

	var devices []Device
	json.Unmarshal(response, &devices)
	return devices
}
