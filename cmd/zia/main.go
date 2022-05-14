package main

import (
	"flag"
	"fmt"
	"os"

	"zscaler_golang/pkg/zia/admin"
	"zscaler_golang/pkg/zia/auth"
	"zscaler_golang/pkg/zia/config"
	"zscaler_golang/pkg/zia/firewall"
	"zscaler_golang/pkg/zia/network"
	"zscaler_golang/pkg/zia/urlcategory"
	"zscaler_golang/pkg/zia/urlfiltering"
)

const (
	name     = "zia_golang"
	version  = "0.0.1"
	revision = "HEAD"
)

func usage() {
	fmt.Fprint(os.Stderr, `Commands:
	zia credential                      # Show credential info placed in config.ini

	zia auth exclude                    # Show exempted urls of cookie auth login

	zia adminuser COMMAND               # Run a command about adminusers
	                                      ls
										  create

	zia adminrole COMMAND               # Run a command about adminusers
	                                      ls

	zia urlcategory COMMAND             # Run a command about urlcategory
	                                      ls
	                                      lookup [URLS

	zia urlfilter COMMAND                 # Run a command about firewall
	                                      ls
	                                     	--id [ID]
	                                     	--all

	zia firewall COMMAND                 # Run a command about firewall
	                                      ls
	                                     	--id [ID]
	                                     	--all

	zia network COMMAND                 # Run a command about network
										  ipdst ls
										  ipsrc ls
`)
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
	case "auth":
		if len(cfg.Args) < 2 {
			fmt.Fprint(os.Stderr, "adminuser: Please specify sub command")
			os.Exit(0)
		}
		if cfg.Args[1] == "exclude" {
			fmt.Printf("%+v\n", auth.FetchExemptedUrls())
		}
	case "adminuser":
		if len(cfg.Args) < 2 {
			fmt.Fprint(os.Stderr, "adminuser: Please specify sub command")
			os.Exit(0)
		}
		if cfg.Args[1] == "ls" {
			fmt.Printf("%+v\n", admin.FetchAllAdminUsers())
		}
		if cfg.Args[1] == "create" {
			if len(cfg.Args) < 3 {
				fmt.Print(`Not enough args. You need fill out :
  - role name
  - login name(email)
  - user name
  - password
  - admin scope`)
				flag.PrintDefaults()
				os.Exit(0)
			}
			fmt.Print(admin.CreateAdminUsers(
				cfg.Args[2], // "role_name",
				cfg.Args[3], // "login_name",
				cfg.Args[4], // "user_name",
				cfg.Args[3], // "email",
				true,        // isPassword bool,
				cfg.Args[5], // "password",
				cfg.Args[6], // "admin_scope",
				"",          // "name",
			))
		}
	case "adminrole":
		if len(cfg.Args) < 2 {
			fmt.Fprint(os.Stderr, "adminrole: Please specify sub command")
			os.Exit(0)
		}
		if cfg.Args[1] == "ls" {
			fmt.Printf("%+v\n", admin.FetchAllAdminRoles())
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
			fmt.Printf("%+v\n", urlcategory.LookupUrlCategory(cfg.Args[2:]))
		} else if cfg.Args[1] == "ls" {
			fmt.Printf("%+v\n", urlcategory.FetchAllUrlCategories())
		}
	case "urlfilter":
		if len(cfg.Args) < 2 {
			fmt.Fprint(os.Stderr, "urlfilter: Please specify sub command")
			os.Exit(0)
		}
		if cfg.Args[1] == "ls" && len(cfg.Args) == 2 {
			fmt.Printf("%+v\n", urlfiltering.FetchAllUrlFilteringRules())
		} else if cfg.Args[1] == "ls" && cfg.Args[2] == "--id" {
			fmt.Printf("%+v\n", urlfiltering.FetchSpecifiedUrlFilteringRule(cfg.Args[3]))
		}
	case "firewall":
		if len(cfg.Args) < 2 {
			fmt.Fprint(os.Stderr, "firewall: Please specify sub command")
			os.Exit(0)
		}
		if cfg.Args[1] == "ls" {
			if len(cfg.Args) > 3 && cfg.Args[2] == "--id" {
				fmt.Printf("%+v\n", firewall.FetchSpecificFwRule(cfg.Args[3]))
			} else if len(cfg.Args) > 2 && cfg.Args[2] == "--all" {
				fmt.Printf("%+v\n", firewall.FetchAllFwRules())
			}
		}
	case "network":
		if len(cfg.Args) < 2 {
			fmt.Fprint(os.Stderr, "network: Please specify sub command")
			os.Exit(0)
		}
		if cfg.Args[1] == "ipdst" {
			if len(cfg.Args) > 2 && cfg.Args[2] == "ls" {
				fmt.Printf("%+v\n", network.FetchIpDstGroups())
			}
		} else if cfg.Args[1] == "ipsrc" {
			if len(cfg.Args) > 2 && cfg.Args[2] == "ls" {
				fmt.Printf("%+v\n", network.FetchIpSrcGroups())
			}
		}
	}
}

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
