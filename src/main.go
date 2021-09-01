package main

import (
	"fmt"
	"zscaler_golang/lib"
)

func main() {
	url_categories := lib.FetchAllUrlCategories()
	fmt.Print(url_categories)
	// target_urls := []string["aaa.com", "bbb.com"]
	// target_urls := []string{"aaa.com", "bbb.com"}
	// category := lib.LookupUrlCategory(target_urls)
	// fmt.Print(category)
}
