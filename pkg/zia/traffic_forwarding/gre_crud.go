package trafficforwarding

import (
	"fmt"
	"net/url"
	"zscaler_golang/pkg/infra"
	"zscaler_golang/pkg/zia/auth"
	"zscaler_golang/pkg/zia/config"
)

type GreTunnel struct {
	Id             int         `json:"id"`
	SourceIp       string      `json:"sourceIp"`
	PrimaryDestVip interface{} `json:"primaryDestIp"`
}

type Vip struct{}

func FetchGreTunnels() {
	sessionId := auth.Login()
	baseUrl, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/greTunnels")
	endpoint := baseUrl.ResolveReference(reference).String()
	response := infra.GetApi(endpoint, sessionId)
	auth.Logout()
	fmt.Print(string(response))
}

func FetchGreTunnelAvailabbleInternalRanges() {
	sessionId := auth.Login()
	baseUrl, _ := url.Parse("https://" + config.Config.Hostname)
	reference, _ := url.Parse("/api/v1/greTunnels/availableInternalIpRanges")
	endpoint := baseUrl.ResolveReference(reference).String()
	response := infra.GetApi(endpoint, sessionId)
	auth.Logout()
	fmt.Print(string(response))
}
