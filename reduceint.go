package main

import "fmt"

func ReduceInt(a []int, f func(int, int) int) {
	result := a[0]
	for _, value := range a[1:] {
		result = f(result, value)
	}
	fmt.Println(result)
}

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}

	ReduceInt(as, mul) // Output: 1000
	ReduceInt(as, sum) // Output: 502
	ReduceInt(as, div) // Output: 250
}
