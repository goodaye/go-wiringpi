package wiringpi

type Pin struct {
	gpio *GPIO
	Code int
}

// Input: Set pin as Input
func (pin Pin) Input() {
	pin.gpio.PinMode(pin.Code, Input)
}

// Output: Set pin as Output
func (pin Pin) Output() {
	pin.gpio.PinMode(pin.Code, Output)
}

// Clock: Set pin as Clock
func (pin Pin) Clock() {
	pin.gpio.PinMode(pin.Code, GpioClock)
}

// Pwm: Set pin as Pwm
func (pin Pin) PWM() {
	pin.gpio.PinMode(pin.Code, PWMOutput)
}

// High: Set pin High
func (pin Pin) High() {
	pin.gpio.DigitalWrite(pin.Code, High)
}

// Low: Set pin Low
func (pin Pin) Low() {
	pin.gpio.DigitalWrite(pin.Code, Low)
}

// // Toggle pin state
// func (pin Pin) Toggle() {
// 	pin.gpio.DigitalWrite(pin.Code, Low)
// }

// Read pin state (high/low)
func (pin Pin) Read() DigitalValue {
	return pin.gpio.DigitalRead(pin.Code)
}
