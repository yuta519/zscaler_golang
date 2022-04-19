package main

import (
	"fmt"
	"zscaler_golang/pkg"
)

func main() {
	// Create New Admin Users
	// adminusers := pkg.CreateAdminUsers(
	// 	"Super Admin",
	// 	"yuta.kawamura@zscaler.net",
	// 	"yuta.kawamura@zscaler.net",
	// 	"yuta.kawamura@zscaler.net",
	// 	true,
	// 	"P@ssw0rd",
	// 	"ORGANIZATION",
	// 	"Yuta Kawamura",
	// )
	// fmt.Print(adminusers)

	// Show Admin Users
	// fmt.Println(pkg.FetchAllNetworkServices())

	// Show all url categories name
	// urlcategories := pkg.FetchAllUrlCategories()
	// fmt.Print(urlcategories)

	// Show url filtering policies
	// pkg.FetchAllUrlFilteringRules()

	// Show url lookup information
	target_urls := []string{"aaa.com", "softbank.com"}
	category := pkg.LookupUrlCategory(target_urls)
	fmt.Print(category)

	// Create New URL Filtering Rule
	// pkg.CreateUrlFilteringRule(
	// 	"READ_WRITE",
	// 	"TEST KAWAMURA RULE",
	// 	1,
	// 	[]string{
	// 		"DOHTTPS_RULE",
	// 		"TUNNELSSL_RULE",
	// 		"HTTP_PROXY",
	// 		"FOHTTP_RULE",
	// 		"FTP_RULE",
	// 		"HTTPS_RULE",
	// 		"HTTP_RULE",
	// 		"SSL_RULE",
	// 		"TUNNEL_RULE",
	// 	},
	// 	nil,
	// 	[]string{"CUSTOM_01", "CUSTOM_02"},
	// 	"ENABLED",
	// 	0,
	// 	[]string{
	// 		"OPTIONS",
	// 		"GET",
	// 		"HEAD",
	// 		"POST",
	// 		"PUT",
	// 		"DELETE",
	// 		"TRACE",
	// 		"CONNECT",
	// 		"OTHER",
	// 	},
	// 	false,
	// 	false,
	// 	0,
	// 	"ALLOW",
	// )

}
