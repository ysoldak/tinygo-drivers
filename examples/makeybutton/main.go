package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/makeybutton"
)

var (
	led    machine.Pin = machine.LED
	button machine.Pin = machine.BUTTON
	key    *makeybutton.Button
)

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	key = makeybutton.NewButton(button)

	for {
		switch key.Get() {
		case makeybutton.Pressed:
			led.High()
		case makeybutton.Released:
			led.Low()
		}
		time.Sleep(30 * time.Millisecond)
	}
}
