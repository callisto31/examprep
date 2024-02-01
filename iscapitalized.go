package main

import (
	"fmt"
)

func isUpperCase(char byte) bool {
	return char >= 'A' && char <= 'Z'
}

func isLowerCase(char byte) bool {
	return char >= 'a' && char <= 'z'
}

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}

	startOfWord := true

	for i := 0; i < len(s); i++ {
		char := s[i]

		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
			if startOfWord && isLowerCase(char) {
				return false
			}
			startOfWord = false
		} else {
			startOfWord = true
		}
        if char >= '0' && char <= '9'{
            return true
        }
    }
    return true
}

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}
