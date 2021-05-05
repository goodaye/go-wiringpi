package  wiringpi

type Pin uint8

// Input: Set pin as Input
func (pin Pin) Input() {
	PinMode(pin, Input)
}

// Output: Set pin as Output
func (pin Pin) Output() {
	PinMode(pin, Output)
}

// Clock: Set pin as Clock
func (pin Pin) Clock() {
	PinMode(pin, Clock)
}

// Pwm: Set pin as Pwm
func (pin Pin) Pwm() {
	PinMode(pin, Pwm)
}

// High: Set pin High
func (pin Pin) High() {
	WritePin(pin, High)
}

// Low: Set pin Low
func (pin Pin) Low() {
	WritePin(pin, Low)
}

// Toggle pin state
func (pin Pin) Toggle() {
	TogglePin(pin)
}