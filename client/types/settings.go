package types

type Overlay string

const (
	OverlayClear   Overlay = "clear"
	OverlaySnow    Overlay = "snow"
	OverlayRain    Overlay = "rain"
	OverlayDrizzle Overlay = "drizzle"
	OverlayStorm   Overlay = "storm"
	OverlayThunder Overlay = "thunder"
	OverlayFrost   Overlay = "frost"
)

type TransitionEffect int

const (
	TransitionEffectRandom   TransitionEffect = iota
	TransitionEffectSlide    TransitionEffect = iota
	TransitionEffectDim      TransitionEffect = iota
	TransitionEffectZoom     TransitionEffect = iota
	TransitionEffectRotate   TransitionEffect = iota
	TransitionEffectPixelate TransitionEffect = iota
	TransitionEffectCurtain  TransitionEffect = iota
	TransitionEffectRipple   TransitionEffect = iota
	TransitionEffectBlink    TransitionEffect = iota
	TransitionEffectReload   TransitionEffect = iota
	TransitionEffectFade     TransitionEffect = iota
)

type Settings struct {
	AppDisplayTime        *int     `json:"ATIME,omitempty"`
	AppStyle              *int     `json:"TMODE,omitempty"`
	AutomaticBrightness   *bool    `json:"ABRI,omitempty"`
	AutomaticSwitchApp    *bool    `json:"ATRANS,omitempty"`
	BatteryApp            *bool    `json:"BAT,omitempty"`
	BatteryColor          *string  `json:"BAT_COL,omitempty"`
	BLOCKN                *bool    `json:"BLOCKN,omitempty"`
	Brightness            *int     `json:"BRI,omitempty"`
	CalendarBodyColor     *int     `json:"CTCOL,omitempty"`
	CalendarHeaderColor   *int     `json:"CHCOL,omitempty"`
	CalendarTextColor     *int     `json:"CBCOL,omitempty"`
	CEL                   *bool    `json:"CEL,omitempty"`
	ColorCorrection       *string  `json:"CCORRECTION,omitempty"`
	ColorTemperature      *string  `json:"CTEMP,omitempty"`
	DateApp               *bool    `json:"DAT,omitempty"`
	DateColor             *string  `json:"DATE_COL,omitempty"`
	DateFormat            *string  `json:"DFORMAT,omitempty"`
	Gamma                 *float64 `json:"GAMMA,omitempty"`
	GlobalTextColor       *int     `json:"TCOL,omitempty"`
	HumidityApp           *bool    `json:"HUM,omitempty"`
	HumidityColor         *string  `json:"HUM_COL,omitempty"`
	MAT                   *int     `json:"MAT,omitempty"`
	MatrixPower           *bool    `json:"MATP,omitempty"`
	Overlay               *Overlay `json:"OVERLAY,omitempty"`
	ScrollSpeed           *int     `json:"SSPEED"`
	SOUND                 *bool    `json:"SOUND,omitempty"`
	StartWeekOnMonday     *bool    `json:"SOM,omitempty"`
	TemperatureApp        *bool    `json:"TEMP,omitempty"`
	TemperatureColor      *string  `json:"TEMP_COL,omitempty"`
	TimeApp               *bool    `json:"TIM,omitempty"`
	TimeColor             *string  `json:"TIME_COL,omitempty"`
	TimeFormat            *string  `json:"TFORMAT,omitempty"`
	TransitionEffect      *int     `json:"TEFF,omitempty"`
	TransitionSpeedMillis *int     `json:"TSPEED,omitempty"`
	Uppercase             *bool    `json:"UPPERCASE,omitempty"`
	Volume                *int     `json:"VOL,omitempty"`
	WeekDay               *bool    `json:"WD,omitempty"`
	WeekDayColorActive    *int     `json:"WDCA,omitempty"`
	WeekDayColorInactive  *int     `json:"WDCI,omitempty"`
}
