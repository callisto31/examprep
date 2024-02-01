package main

import (
    "fmt"
    "os"
)

func check(str1, str2 string) bool {
    istr2 := 0
    for _, char := range str1 {
        index := -1
        for i := istr2; i < len(str2); i++ {
            if rune(str2[i]) == char {
                index = i - istr2
                break
            }
        }
        if index == -1 {
            return false
        }
        istr2 += index + 1
    }
    return true
}

func main() {
    if len(os.Args) != 3 {
        return
    }
    str1, str2 := os.Args[1], os.Args[2]
    if check(str1, str2) {
        fmt.Println(str1)
    }
}
