package main

import (
	"fmt"
	"zscaler_golang/pkg"
)

func main() {
	// pkg.FetchAllUrlCategories()
	fmt.Print(pkg.FetchAllUrlFilteringRules())
	// fmt.Print(pkg.FetchAllUrlCategories())
	// target_urls := []string{"aaa.com", "bbb.com"}
	// category := pkg.LookupUrlCategory(target_urls)
	// fmt.Print(category)
}
