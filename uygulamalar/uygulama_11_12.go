// uygulama_11_12
// https://gobot.io/documentation/platforms/raspi/
package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
	"time"
)

func main() {
	gbot := gobot.NewGobot()
	r := raspi.NewRaspiAdaptor("raspi")
	led := gpio.NewLedDriver(r, "led", "2")
	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}
	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{led},
		work,
	)
	gbot.AddRobot(robot)
	gbot.Start()
}
