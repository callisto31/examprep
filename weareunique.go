package piscine

func WeAreUnique(str1, str2 string) int {
    str3 := str1 + str2
    var count int
    if str3 == "" {
        return -1
    }
    mmap := make(map[rune]int)
    for _, l := range str3 {
        mmap[l]++
    }
    for _, mpl := range mmap {
        if mpl == 1 {
            count++
        } else if mpl < 2 {
            return 0
        }
    }
    return count
}
