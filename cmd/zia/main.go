package main

import (
	"fmt"
	"os"

	"zscaler_golang/pkg/zia/admin"
	"zscaler_golang/pkg/zia/config"
)

const (
	name     = "zia_golang"
	version  = "0.0.1"
	revision = "HEAD"
)

func usage() {
	// 	fmt.Fprint(os.Stderr, `Tasks:
	//   goreman check                      # Show entries in Procfile
	//   goreman help [TASK]                # Show this help
	//   goreman export [FORMAT] [LOCATION] # Export the apps to another process
	//                                        (upstart)
	//   goreman run COMMAND [PROCESS...]   # Run a command
	//                                        start
	//                                        stop
	//                                        stop-all
	//                                        restart
	//                                        restart-all
	//                                        list
	//                                        status
	//   goreman start [PROCESS]            # Start the application
	//   goreman version                    # Display Goreman version
	// Options:
	// `)
	// 	flag.PrintDefaults()
	os.Exit(0)
}

func readConfg() {
	fmt.Print(config.Config.ApiKey)
}

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
	// fmt.Println(admin.FetchAllAdminRoles())
	// readConfg()
	fmt.Println(admin.FetchAllAdminUsers())

	// Show all url categories name
	// urlcategories := pkg.FetchAllUrlCategories()
	// fmt.Print(urlcategories)

	// Show url filtering policies
	// pkg.FetchAllUrlFilteringRules()

	// Show url lookup information
	// target_urls := []string{"aaa.com", "softbank.com"}

	// category := urlcategory.LookupUrlCategory(target_urls)
	// fmt.Print(category)

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
