package gumeng

import (
	"net/url"
	"sort"
	"strings"

	"github.com/usthooz/gutil"
)

// Sign 请求签名
func Sign(apiSecurity string, path string, params url.Values) string {
	var (
		keys []string
	)
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		path += key + params.Get(key)
	}
	sig := gutil.HmacSHA1(apiSecurity, path)
	return strings.ToUpper(sig)
}
