package wiringpi

// #cgo LDFLAGS: -lwiringPi
// #include <wiringPi.h>
import "C"

var loaded = false

// PinMode enumerates available pin modes
type PinMode int

const (
	// Input set pin for input
	Input PinMode = C.INPUT
	// Output set pin for output
	Output PinMode = C.OUTPUT
	// PWMOutput set pin for PWM output
	PWMOutput PinMode = C.PWM_OUTPUT
	// GpioClock set pin for clock mode
	GpioClock PinMode = C.GPIO_CLOCK
)

// PullMode enumerates available pull modes
type PullMode int

const (
	// PullOff turns off pull up/down
	PullOff PullMode = C.PUD_OFF
	// PullDown pulls pin down
	PullDown PullMode = C.PUD_DOWN
	// PullUp pulls pin up
	PullUp PullMode = C.PUD_UP
)

// DigitalValue enumberates digital IO values
type DigitalValue int

const (
	// High pin is at high state
	High DigitalValue = C.HIGH
	// Low pin is at low state
	Low DigitalValue = C.LOW
)

// SetPinMode sets pin mode
func SetPinMode(pin int, mode PinMode) {
	C.pinMode(C.int(pin), C.int(mode))
}

// Pull sets pull up/down mode
func Pull(pin int, mode PullMode) {
	C.pullUpDnControl(C.int(pin), C.int(mode))
}

// DigitalWrite writes digital value to pin
func DigitalWrite(pin int, val DigitalValue) {
	C.digitalWrite(C.int(pin), C.int(val))
}

// PWMSetRange set PWM generator range
// defaults to 1024
func PWMSetRange(val uint) {
	C.pwmSetRange(C.uint(val))
}

// PWMSetClock sets the divisor of PWM clock
func PWMSetClock(val int) {
	C.pwmSetClock(C.int(val))
}

// PWMWrite writes pwn value
func PWMWrite(pin int, val int) {
	C.pwmWrite(C.int(pin), C.int(val))
}

// DigitalRead reads digital value
func DigitalRead(pin int) DigitalValue {
	return DigitalValue(C.digitalRead(C.int(pin)))
}

// Setup setup the GPIO interface
func Setup(method SetupMethod) error {
	if loaded {
		panic("wiring pi is already loaded")
	}
	var ret C.int
	switch method {
	case WiringPiSetup:
		ret = C.wiringPiSetup()
	case BroadcomSetup:
		ret = C.wiringPiSetupGpio()
	case PhysSetup:
		ret = C.wiringPiSetupPhys()
	case SysSetup:
		ret = C.wiringPiSetupSys()
	}
	if ret != 0 {
		return RetCode{int(ret)}
	}
	loaded = true
	return nil
}
