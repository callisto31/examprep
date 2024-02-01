package main

func ItoaBase(value, base int) string {
	if value == 0 {
		return "0"
	}
	digits := "0123456789ABCDEF"
	isNegative := false
	if value < 0 {
		isNegative = true
		value = -value
	}
	var result string
	for value > 0 {
		remainder := value % base
		result = string(digits[remainder]) + result
		value /= base
	}
	if isNegative {
		result = "-" + result
	}
	return result
}
func main() {
	// Test the function
	result := ItoaBase(255, 16)
	println(result)
}
