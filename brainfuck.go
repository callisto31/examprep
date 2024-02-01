package main
import (
	"fmt"
	"os"
)
const SIZE = 2048
func main() {
	if len(os.Args) != 2 {
		return
	}
	progpoint := []byte(os.Args[1])
	var arby [SIZE]byte
	pos := 0
	openBr := 0         
	i := 0              
	N := len(progpoint) 
	for i >= 0 && i < N {
		switch progpoint[i] {
		case '>':
			pos++
		case '<':
			pos--
		case '+':
			arby[pos]++
		case '-':
			arby[pos]--
		case '.':
			fmt.Printf("%c", rune(arby[pos]))
		case '[':
			openBr = 0
			if arby[pos] == 0 {
				for i < N && (progpoint[i] != byte(']') || openBr > 1) {
					if progpoint[i] == byte('[') {
						openBr++
					} else if progpoint[i] == byte(']') {
						openBr--
					}
					i++
				}
			}
		case ']':
			openBr = 0
			if arby[pos] != 0 {
				for i >= 0 && (progpoint[i] != byte('[') || openBr > 1) {
					if progpoint[i] == byte(']') {
						openBr++
					} else if progpoint[i] == byte('[') {
						openBr--
					}
					i--
				}
			}
		}
		i++
	}
}
