package main

import (
	"os"
	"strings"
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

	for _, char := range str1 {
		if !seen[char] && strings.ContainsRune(str2, char) {
			result += string(char)
			seen[char] = true
		}
	}

	for _, char := range result {
		z01.PrintRune(char)
	}

	z01.PrintRune('\n')
}
