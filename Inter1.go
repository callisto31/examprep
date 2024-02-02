package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	result := ""
	seen := make(map[rune]bool)

	for _, char1 := range str1 {
		if !seen[char1] {
			for _, char2 := range str2 {
				if char1 == char2 {
					result += string(char1)
					seen[char1] = true
					break
				}
			}
		}
	}

	for _, char := range result {
		z01.PrintRune(char)
	}

	z01.PrintRune('\n')
}
