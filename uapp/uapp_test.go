package uapp

import (
	"testing"
)

// TestEventKeyUV
func TestEventKeyUV(t *testing.T) {
	uapp := &Uapp{
		ApiKey:      "",
		ApiSecurity: "",
		GateWay:     "https://gateway.open.umeng.com",
		AppKey:      "",
		Debug:       true,
	}
	newData, err := uapp.GetNewUserData("2020-12-28", "2020-12-28")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp new-> %#v", newData)

	activeData, err := uapp.GetActiveUserData("2020-12-28", "2020-12-28")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp active-> %#v", activeData)

	eventData, err := uapp.GetEventKeyUV("com.app.home", "2020-12-28", "2020-12-28")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp event-> %#v", eventData)
}
