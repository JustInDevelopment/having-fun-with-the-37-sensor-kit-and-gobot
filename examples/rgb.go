package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("COM112")
	ledGreen := gpio.NewLedDriver(firmataAdaptor, "6")
	ledBlue := gpio.NewLedDriver(firmataAdaptor, "5")
	ledRed := gpio.NewLedDriver(firmataAdaptor, "3")
	ledCount := 0
	rgb := ledBlue

	work := func() {

		for light := 0; light <= 255; light++ {
			time.Sleep(time.Millisecond * 300)
			for light := 255; light >= 0; light-- {

				light := byte(light)
				rgb.Brightness(light)
			}
			light := byte(light)
			rgb.Brightness(light)
			if ledCount == 0 {
				rgb = ledGreen
				ledBlue.Off()
				ledRed.Off()
			}

			if ledCount == 1 {
				rgb = ledBlue
				ledGreen.Off()
				ledRed.Off()
			}

			if ledCount == 2 {
				rgb = ledRed
				ledBlue.Off()
				ledGreen.Off()
			}

			ledCount++
			if ledCount == 3 {
				ledCount = 0
			}

		}

	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{ledGreen, ledBlue, ledRed},
		work,
	)

	robot.Start()

}
