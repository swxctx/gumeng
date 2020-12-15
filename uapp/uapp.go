package uapp

// Uapp
type Uapp struct {
	// ApiKey
	ApiKey string
	// ApiSecurity
	ApiSecurity string
	// GateWay
	GateWay string
	// AppKey
	AppKey string
	// Debug 调试标识
	Debug bool
}

// NewUapp
func NewUapp(apiKey, apiSecurity, gateWay string, debug bool) *Uapp {
	return &Uapp{
		ApiKey:      apiKey,
		ApiSecurity: apiSecurity,
		GateWay:     gateWay,
		Debug:       debug,
	}
}
