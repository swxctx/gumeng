package uapp

import (
	"testing"
)

// TestEventKeyUV
func TestEventKeyUV(t *testing.T) {
	uapp := &Uapp{
		ApiKey:      "2558925",
		ApiSecurity: "Pn0Zf6vR2wmK",
		GateWay:     "https://gateway.open.umeng.com",
		AppKey:      "",
		Debug:       true,
	}
	eventUV, err := uapp.GetEventKeyUV("com.swxctx.com", "2020-12-14", "2020-12-15")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp-> %v", eventUV)
}
