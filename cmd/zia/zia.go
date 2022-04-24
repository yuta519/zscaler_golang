package main

import (
	"flag"
	"fmt"
	"os"

	"zscaler_golang/pkg/zia/admin"
	"zscaler_golang/pkg/zia/config"
	"zscaler_golang/pkg/zia/urlcategory"
)

const (
	name     = "zia_golang"
	version  = "0.0.1"
	revision = "HEAD"
)

func usage() {
	fmt.Fprint(os.Stderr, `Tasks:
	zia adminuser [TASK]                # Show adminusers in Zscaler Internet Access
	zia adminrole [TASK]                # Show adminroles in Zscaler Internet Access
Options:
`)

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
	flag.PrintDefaults()
	os.Exit(0)
}

type conf struct {
	UserName string
	Password string
	Hostname string
	ApiKey   string
	Args     []string
}

func readConfg() *conf {
	var cfg conf

	flag.Parse()
	if flag.NArg() == 0 {
		usage()
	}

	cfg.UserName = config.Config.UserName
	cfg.Password = config.Config.Password
	cfg.Hostname = config.Config.Hostname
	cfg.ApiKey = config.Config.ApiKey
	cfg.Args = flag.Args()

	return &cfg
}

func main() {
	var err error
	if err != nil {
		fmt.Fprint(os.Stderr, "zia: ", err.Error())
	}
	cfg := readConfg()

	cmd := cfg.Args[0]
	switch cmd {
	case "credential":
		fmt.Print("username: ", cfg.UserName, "\n")
		fmt.Print("cloud console: ", cfg.Hostname, "\n")
		fmt.Print("password: *****", cfg.Password[6:], "\n")
		fmt.Print("apikey: *****", cfg.ApiKey[8:], "\n")
	case "adminuser":
		if len(cfg.Args) < 2 {
			fmt.Fprint(os.Stderr, "adminuser: Please specify sub command")
			os.Exit(0)
		}
		if cfg.Args[1] == "ls" {
			fmt.Println(admin.FetchAllAdminUsers())
		}
	case "urlcategory":
		if len(cfg.Args) < 2 {
			fmt.Fprint(os.Stderr, "urlcategory: Please specify sub command")
			os.Exit(0)
		}
		if cfg.Args[1] == "lookup" {
			if len(cfg.Args) < 3 {
				fmt.Fprint(os.Stderr, "lookup: Please input urls")
				os.Exit(0)
			}
			fmt.Println(urlcategory.LookupUrlCategory(cfg.Args[2:]))
		} else if cfg.Args[1] == "ls" {
			fmt.Println(urlcategory.FetchAllUrlCategories())
		}
	}
}

// fmt.Print(cmd)
// fmt.Print(a)
// fmt.Print(&a)
// fmt.Print(*a)
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
// fmt.Println(admin.FetchAllAdminUsers())

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
