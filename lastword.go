package main

import (
    "fmt"
    "os"
)

func lastword(s string) {
    last := ""
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == ' ' {
            if last != "" {
                break
            }
        }
        last = string(s[i]) + last
    }
    fmt.Println(last)
}

func main() {
    if len(os.Args) == 2 {
        input := os.Args[1]
        lastword(input)
    }
}
