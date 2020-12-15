package gumeng

import (
	"net/url"
	"testing"
)

// TestSign
func TestSign(t *testing.T) {
	params := url.Values{}
	params.Add("b", "2")
	params.Add("a", "1")
	t.Logf("Sign-> %s", Sign("1000000", "param2/1/system/currentTime/1000000", params))
}
