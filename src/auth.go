package main

import (
	"fmt"
	"time"
)

type Auth struct {
	ApiKey string
	Now    string
}

func (a *Auth) obfuscateApiKey() {
	unix_now := time.Now().Unix()
	fmt.Println(unix_now)
}

func main() {
	auth := Auth{"aaaaa", "11111111"}
	auth.obfuscateApiKey()
}
