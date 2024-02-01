package main

import "fmt"

func Fishandchips(n int32) string {
	switch {
	case n%2 == 0 && n%3 == 0:
		return "fish and chips$\n"
	case n%3 == 0:
		return "chips$\n"
	case n%2 == 0:
		return "$\n"
	default:
		return "$\n"
	}
}

func main() {
	args := []int32{0, 6, 33, -3, 5, 8}
	for i := 0; i < len(args); i++ {
		fmt.Print(Fishandchips(args[i]))
	}
}
