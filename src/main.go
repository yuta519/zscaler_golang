package main

import (
	"encoding/json"
	"fmt"
	"zscaler_golang/sdk"
)

func main() {
	url_categories := sdk.FetchAllUrlCategories()
	var response [][]string
	if err := json.Unmarshal([]byte(url_categories), &response); err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(url_categories))
	// fmt.Println(url_categories)
	fmt.Printf("%s", url_categories)
	// for _, url_category := range url_categories {
	// 	fmt.Printf("%s", url_category)
	// }
	// target_urls := []string["aaa.com", "bbb.com"]
	// target_urls := []string{"aaa.com", "bbb.com"}
	// category := lib.LookupUrlCategory(target_urls)
	// fmt.Print(category)
}
