package main

import "fmt"

func FromTo(from, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid$\n"
	}

	var result string
	if from <= to {
		for i := from; i <= to; i++ {
			if i < 10 {
				result += "0" + fmt.Sprint(i) + ", "
			} else {
				result += fmt.Sprint(i) + ", "
			}
		}
	} else {
		for i := from; i >= to; i-- {
			if i < 10 {
				result += "0" + fmt.Sprint(i) + ", "
			} else {
				result += fmt.Sprint(i) + ", "
			}
		}
	}

	if len(result) > 2 {
		result = result[:len(result)-2] + "$\n"
	} else {
		result += "$\n"
	}

	return result
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}
