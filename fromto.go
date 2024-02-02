package main

import (
    "fmt"
)

func FromTo(from int, to int) string {
    if from < 0 || from > 99 || to < 0 || to > 99 {
        return "Invalid\n"
    }

    var result string
    for i := from; i <= to; i++ {
        if i < 10 {
            result += fmt.Sprintf("0%d, ", i)
        } else {
            result += fmt.Sprintf("%02d, ", i)
        }
    }

    return result[:len(result)-2] + "\n"
}

func main() {
    fmt.Print(FromTo(1, 10))
    fmt.Print(FromTo(10, 1))
    fmt.Print(FromTo(10, 10))
    fmt.Print(FromTo(100, 10))
}
