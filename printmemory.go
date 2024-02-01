package main

import "fmt"

func PrintMemory(arr [10]byte) {
    arr2 := arr
    base := [16]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
    var outer [][]int
    var inner []int
    for i := len(arr) - 1; i >= 0; i-- {
        if arr[i] >= 0 && arr[i] < 16 {
            inner = append(inner, int(arr[i]), 0)
            outer = append(outer, inner)
            inner = nil
            continue
        }
        for arr[i] != 0 {
            inner = append(inner, int(arr[i]%16))
            arr[i] = arr[i] / 16
        }
        outer = append(outer, inner)
        inner = nil
    }

    for i := len(outer) - 1; i >= 0; i-- {
        if i == 0 || i == 2 || i == 3 || i == 4 || i == 6 || i == 7 || i == 8 {
            fmt.Print(" ")
        }
        if i == 5 || i == 1 {
            fmt.Print("\n")
        }
        for j := len(outer[i]) - 1; j >= 0; j-- {
            fmt.Print(base[outer[i][j]])
        }
    }
    fmt.Print("\n")
    for _, dec := range arr2 {
        if dec >= 0 && dec <= 31 {
            fmt.Print(".")
        } else {
            fmt.Print(string(dec))
        }
    }
    fmt.Print("\n")
}
func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
