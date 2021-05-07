package wiringpi

type Pin int

// Input: Set pin as Input
func (pin Pin) Input() {
	SetPinMode(int(pin), Input)
}

// Output: Set pin as Output
func (pin Pin) Output() {
	SetPinMode(int(pin), Output)
}

// Clock: Set pin as Clock
func (pin Pin) Clock() {
	SetPinMode(int(pin), GpioClock)
}

// Pwm: Set pin as Pwm
func (pin Pin) PWM() {
	SetPinMode(int(pin), PWMOutput)

}

// High: Set pin High
func (pin Pin) High() {
	DigitalWrite(int(pin), High)
}

// Low: Set pin Low
func (pin Pin) Low() {
	DigitalWrite(int(pin), Low)

}

// // Toggle pin state
// func (pin Pin) Toggle() {
// 	pin.gpio.DigitalWrite(pin.Code, Low)
// }

// Read pin state (high/low)
func (pin Pin) Read() DigitalValue {
	return DigitalRead(int(pin))
}
