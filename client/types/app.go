package types

type LifetimeMode int

const (
	LifetimeModeDelete LifetimeMode = iota
	LifetimeModeStale  LifetimeMode = iota
)

type TextCase int

const (
	TextCaseSystem    TextCase = iota
	TextCaseUppercase TextCase = iota
	TextCaseAsSent    TextCase = iota
)

type PushIcon int

const (
	PushIconFixed    PushIcon = iota
	PushIconMoveOnce PushIcon = iota
	PushIconMove     PushIcon = iota
)

// App /api/custom POST
type App struct {
	// The text to display.Keep in mind the font does not have a fixed size and I uses less space than W.This facts affects when text will start scrolling
	Text *string `json:"text,omitempty"`
	// Changes the Uppercase setting.0 = global setting, 1 = forces uppercase 2 = shows as it sent
	TextCase *TextCase `json:"textCase,omitempty"`
	// Draw the text on top.false
	TopText *bool `json:"topText,omitempty"`
	// Sets an offset for the x position of a starting text.0
	TextOffset *int `json:"textOffset,omitempty"`
	// Centers a short, non-scrollable text.true
	Center *bool `json:"center,omitempty"`
	// The text, bar or line color.
	Color *string `json:"color,omitempty"`
	// Colorizes the text in a gradient of two given colors
	Gradient []string `json:"gradient,omitempty"`
	// Blinks the text in a given interval in ms, not compatible with gradient or rainbow
	BlinkText *int `json:"blinkText,omitempty"`
	// Fades the text on and off in a given interval, not compatible with gradient or rainbow
	FadeText *int `json:"fadeText,omitempty"`
	// Sets a background color.N/A
	Background *string `json:"background,omitempty"`
	// Fades each letter in the text differently through the entire RGB spectrum.false
	Rainbow *bool `json:"rainbow,omitempty"`
	// The icon ID or filename (without extension) to display on the app.You can also send a 8x8 jpg as Base64 String
	Icon *int `json:"icon,omitempty"`
	// 0 = Icon doesn't move. 1 = Icon moves with text and will not appear again. 2 = Icon moves with text but appears again when the text starts to scroll again.
	PushIcon *PushIcon `json:"pushIcon,omitempty"`
	// Sets how many times the text should be scrolled through the matrix before the app ends.
	Repeat *int `json:"repeat,omitempty"`
	// Sets how long the app or notification should be displayed.5
	Duration *int `json:"duration,omitempty"`
	// Set it to true, to hold your notification on top until you press the middle button or dismiss it via HomeAssistant.This key only belongs to notification.
	Hold *bool `json:"hold,omitempty"`
	// The filename of your RTTTL ringtone file placed in the MELODIES folder (without extension).
	Sound *string `json:"sound,omitempty"`
	// Allows to send the RTTTL sound string with the json.
	Rtttl *string `json:"rtttl,omitempty"`
	// Loops the sound or rtttl as long as the notification is running.
	LoopSound *bool `json:"loopSound,omitempty"`
	// Draws a bargraph.Without icon maximum 16 values, with icon 11 values.
	Bar []int `json:"bar,omitempty"`
	// Draws a linechart.Without icon maximum 16 values, with icon 11 values.
	Line []int `json:"line,omitempty"`
	// Enables or disables autoscaling for bar and linechart.
	Autoscale *bool `json:"autoscale,omitempty"`
	// Shows a progress bar.Value can be 0-100.
	Progress *int `json:"progress,omitempty"`
	// The color of the progress bar.-1 X X
	ProgressC *string `json:"progressC,omitempty"`
	// The color of the progress bar background.
	ProgressBC *string `json:"progressBC,omitempty"`
	// Defines the position of your custom page in the loop, starting at 0 for the first position.This will only apply with your first push.This function is experimental.
	Pos *int `json:"pos,omitempty"`
	// Array of drawing instructions.Each object represents a drawing command.See the drawing instructions below.
	Draw map[string][]Drawable `json:"draw,omitempty"`
	// Removes the custom app when there is no update after the given time in seconds.
	Lifetime *int `json:"lifetime,omitempty"`
	// 0 = deletes the app, 1 = marks it as staled with a red rectangle around the app
	LifetimeMode *LifetimeMode `json:"lifetimeMode,omitempty"`
	// Defines if the notification will be stacked.false will immediately replace the current notification.
	Stack *bool `json:"stack,omitempty"`
	// If the Matrix is off, the notification will wake it up for the time of the notification.
	Wakeup *bool `json:"wakeup,omitempty"`
	// Disables the text scrolling.
	NoScroll *bool `json:"noScroll,omitempty"`
	// Allows forwarding a notification to other awtrix devices.Use the MQTT prefix for MQTT and IP addresses for HTTP.
	Clients []string `json:"clients,omitempty"`
	// Modifies the scroll speed.Enter a percentage value of the original scroll speed.
	ScrollSpeed *int `json:"scrollSpeed,omitempty"`
	// Shows an effect as background.The effect can be removed by sending an empty string for effect
	Effect *string `json:"effect,omitempty"`
	// Changes color and speed of the effect
	EffectSettings map[string]string `json:"effectSettings,omitempty"`
	// Saves your custom app into flash and reloads it after boot. Avoid this for custom apps with high update
	// frequencies because the ESP's flash memory has limited write cycles.
	Save *bool `json:"save,omitempty"`
	// Sets an effect overlay (cannot be used with global overlays).
	Overlay *string `json:"overlay,omitempty"`
}
