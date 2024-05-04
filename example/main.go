package main

import (
	"fmt"
	"os"

	"github.com/bakito/go-awtrix/client"
	"github.com/bakito/go-awtrix/client/types"
)

func main() {
	cl := client.New(os.Args[1])
	var err error
	// err = cl.UpdateCustomApp("hello", types.NewAppBuilder().Text("Hello Awtrix").Icon(4123))

	if false {
		l, err := cl.Loop()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v", l)
	}

	_ = cl.Notify(types.NewAppBuilder().
		Duration(10).
		PushIcon(types.PushIconMove).
		TextCase(types.TextCaseAsSent).
		Icon(7270).
		Text("Hello Awtrix").
		Rainbow(true),
	)

	if false {
		err = cl.UpdateSettings(types.NewSettingsBuilder().
			Overlay(types.OverlayClear).
			BatteryApp(true).
			DateApp(true).
			HumidityApp(false).
			TemperatureApp(false),
		)
	}
	if err != nil {
		panic(err)
	}
}
