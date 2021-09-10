package main

import (
	"fmt"
	"zscaler_golang/pkg"
)

func main() {
	urlcategories := pkg.FetchAllUrlCategories()
	fmt.Print(urlcategories)
	// pkg.FetchAllUrlFilteringRules()
	// fmt.Print(pkg.FetchAllUrlCategories())
	// target_urls := []string{"aaa.com", "softbank.com"}
	// category := pkg.LookupUrlCategory(target_urls)
	// fmt.Print(category)
	// fmt.Println(pkg.FetchAllNetworkServices())
}
