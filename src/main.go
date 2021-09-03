package main

import (
	"fmt"
	"zscaler_golang/sdk"
)

func main() {
	sdk.FetchAllUrlCategories()
	target_urls := []string{"aaa.com", "bbb.com"}
	category := sdk.LookupUrlCategory(target_urls)
	fmt.Print(category)
}
