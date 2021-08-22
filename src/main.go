// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"time"

// 	"zscaler_golang/auth"
// 	"zscaler_golang/config"
// )

// type Auth struct {
// 	ApiKey string
// }

// func (a *Auth) obfuscateApiKey() string {
// 	unix_now := time.Now().UnixNano() / int64(time.Millisecond)
// 	convert_str_unix := strconv.FormatInt(unix_now, 10)
// 	key_from_unix := convert_str_unix[len(convert_str_unix)-6:]
// 	r, _ := strconv.Atoi(key_from_unix)
// 	shifted_key := fmt.Sprintf("%06d", r>>1)

// 	apikey := a.ApiKey
// 	var obfuscatedApiKey string

// 	for _, i := range key_from_unix {
// 		index, _ := strconv.Atoi(string(i))
// 		obfuscatedApiKey += string(apikey[index])
// 	}

// 	for _, i := range shifted_key {
// 		index, _ := strconv.Atoi(string(i))
// 		obfuscatedApiKey += string(apikey[index])
// 	}

// 	fmt.Println(obfuscatedApiKey)
// 	return obfuscatedApiKey

// }

// func main() {
// 	auth := Auth{config.Config.ApiKey}
// 	auth.obfuscateApiKey()
// }
