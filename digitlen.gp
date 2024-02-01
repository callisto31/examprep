package main

import "fmt"

func DigitLen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	}
	if n < 0 {
		n *= -1
	}
	if n == 0 {
		return 1
	}
	digitCount := 0
	for n > 0 {
		n = n / base
		digitCount++
	}
	return digitCount
}
func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}
