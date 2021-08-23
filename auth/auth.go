package auth

import (
	"fmt"
	"strconv"
	"time"
	"zscaler_golang/config"
)

type AuthPrepare struct {
	ObfuscatedApiKey string
}

var Auth AuthPrepare

func init() {
	unix_now := time.Now().UnixNano() / int64(time.Millisecond)
	convert_str_unix := strconv.FormatInt(unix_now, 10)
	key_from_unix := convert_str_unix[len(convert_str_unix)-6:]
	r, _ := strconv.Atoi(key_from_unix)
	shifted_key := fmt.Sprintf("%06d", r>>1)

	apikey := config.Config.ApiKey
	var obfuscatedApiKey string

	for _, i := range key_from_unix {
		index, _ := strconv.Atoi(string(i))
		obfuscatedApiKey += string(apikey[index])
	}

	for _, i := range shifted_key {
		index, _ := strconv.Atoi(string(i))
		obfuscatedApiKey += string(apikey[index])
	}

	Auth = AuthPrepare{
		ObfuscatedApiKey: obfuscatedApiKey,
	}

}
