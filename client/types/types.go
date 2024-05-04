package types

// https://blueforcer.github.io/awtrix3/#/api

// https://www.youtube.com/watch?v=AkCocqeyH_U

// DeviceStats /api/stats
type DeviceStats struct {
	Bat        int    `json:"bat,omitempty"`
	BatRaw     int    `json:"bat_raw,omitempty"`
	Type       int    `json:"type,omitempty"`
	Lux        int    `json:"lux,omitempty"`
	LdrRaw     int    `json:"ldr_raw,omitempty"`
	Ram        int    `json:"ram,omitempty"`
	Bri        int    `json:"bri,omitempty"`
	Temp       int    `json:"temp,omitempty"`
	Hum        int    `json:"hum,omitempty"`
	Uptime     int    `json:"uptime,omitempty"`
	WifiSignal int    `json:"wifi_signal,omitempty"`
	Messages   int    `json:"messages,omitempty"`
	Version    string `json:"version,omitempty"`
	Indicator1 bool   `json:"indicator1,omitempty"`
	Indicator2 bool   `json:"indicator2,omitempty"`
	Indicator3 bool   `json:"indicator3,omitempty"`
	App        string `json:"app,omitempty"`
	Uid        string `json:"uid,omitempty"`
	Matrix     bool   `json:"matrix,omitempty"`
	IpAddress  string `json:"ip_address,omitempty"`
}

// Effects /api/effects
type Effects []string

// Transitions /api/transitions
type Transitions []string

// Loop  /api/loop
type Loop map[string]int

//	Time        int `json:"Time,omitempty"`
//	Temperature int `json:"Temperature,omitempty"`
//	Humidity    int `json:"Humidity,omitempty"`
//	Battery     int `json:"Battery,omitempty"`

// Screen /api/screen
type Screen []int32

// Power /api/power POST
type Power struct {
	Power bool `json:"power,omitempty"`
}
