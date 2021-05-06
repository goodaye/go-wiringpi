package main

import (
	"fmt"

	"github.com/goodaye/go-wiringpi"
)

func main() {
	gpio, err := wiringpi.Setup(wiringpi.WiringPiSetup)
	if err != nil {
		panic(err)
	}
	gpio.PinMode(1, wiringpi.Input)
	if gpio.DigitalRead(1) == wiringpi.High {
		fmt.Println("WiringPi port 1 is at high")
	} else {
		fmt.Println("WiringPi port 1 is at low")
	}

	p1 := gpio.Pin(10)
	p1.Output()
	p1.High()
}
