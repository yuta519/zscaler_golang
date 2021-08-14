package main

import (
	"fmt"
	"strconv"
	"time"
)

type Auth struct {
	ApiKey string
	Now    string
}

func (a *Auth) obfuscateApiKey() {
	unix_now := time.Now().UnixNano() / int64(time.Millisecond)
	convert_str_unix := strconv.FormatInt(unix_now, 10)
	key_from_unix := convert_str_unix[len(convert_str_unix)-6:]
	r, _ := strconv.Atoi(key_from_unix)
	shifted_key := fmt.Sprintf("%06d", r>>1)

	// var obfuscateApiKey string
	int_key_from_unix, _ := strconv.Atoi(key_from_unix)
	fmt.Println(int_key_from_unix)
	// for i := 0; i < len(int_key_from_unix); i++ {
	// 	pass
	// 	fmt.Println()
	// }

	fmt.Println(key_from_unix)
	fmt.Println(shifted_key)

}

func main() {
	auth := Auth{"aaaaa", "11111111"}
	auth.obfuscateApiKey()
}
