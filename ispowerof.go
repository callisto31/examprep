package main

import (
    "fmt"
    "os"
    "strconv"
)

func ispower(n int) bool {
    return n > 0 && (n&(n-1)) == 0
}

func main() {
    if len(os.Args) == 2 {
        if num, err := strconv.Atoi(os.Args[1]); err == nil {
            fmt.Println(ispower(num))
        }
    }
}
