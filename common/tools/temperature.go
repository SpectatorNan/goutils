package tools

// Celsius convert to fahrenheit
func TemperatureCelsiusToFahrenheit(celsius float32) float32 {
	return celsius*(9/5.0) + 32
}

// Fahrenheit convert to celsius
func TemperatureFahrenheitToCelsius(fahrenheit float32) float32 {
	result := (fahrenheit - 32) / (9 / 5.0)
	return result
}
