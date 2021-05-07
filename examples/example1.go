package main

import (
	"fmt"

	"github.com/goodaye/go-wiringpi"
)

func main() {
	err := wiringpi.Setup(wiringpi.WiringPiSetup)
	if err != nil {
		panic(err)
	}
	wiringpi.SetPinMode(1, wiringpi.Input)
	if wiringpi.DigitalRead(1) == wiringpi.High {
		fmt.Println("WiringPi port 1 is at high")
	} else {
		fmt.Println("WiringPi port 1 is at low")
	}

	p1 := wiringpi.Pin(10)
	p1.Output()
	p1.High()
}
