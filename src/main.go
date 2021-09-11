package main

import (
	"fmt"
	"zscaler_golang/pkg"
)

func main() {
	// Show Admin Users
	adminusers := pkg.CreateAdminUsers(
		"Super Admin",
		"yuta.kawamura@zscaler.net",
		"yuta.kawamura@zscaler.net",
		"yuta.kawamura@zscaler.net",
		true,
		"P@ssw0rd",
		"ORGANIZATION",
		"Yuta Kawamura",
	)
	fmt.Print(adminusers)

	// Show Admin Users
	// fmt.Println(pkg.FetchAllNetworkServices())

	// Show all url categories name
	// urlcategories := pkg.FetchAllUrlCategories()
	// fmt.Print(urlcategories)

	// Show url filtering policies
	// pkg.FetchAllUrlFilteringRules()

	// Show url lookup information
	// target_urls := []string{"aaa.com", "softbank.com"}
	// category := pkg.LookupUrlCategory(target_urls)
	// fmt.Print(category)
}
