package main

import (
	"fmt"
	"io"
	//"log"
	"os"
)

func main() {
	var ch rune
	for {
		if _, err := fmt.Fscanf(os.Stdin, "%c", &ch); err != nil {
			if err == io.EOF {
				break
			}
		} else {
			switch {
			case ch == 'j':
				fallthrough
			case ch == 'h':
				fmt.Print("I")
			case ch == 'ñ':
				fmt.Print("N")
			case ch == 'k':
				fmt.Print("L")
			case ch == 'u':
				fallthrough
			case ch == 'w':
				fmt.Print("V")
			case ch == 'y':
				fmt.Print("Z")
			case ch == 'á':
				fmt.Print("A")
			case ch == 'é':
				fmt.Print("E")
			case ch == 'í':
				fmt.Print("I")
			case ch == 'ó':
				fmt.Print("O")
			case ch == 'ú':
				fmt.Print("V")
			case 'a' <= ch && ch <= 'z':
				fmt.Printf("%c", ch-32)
			default:
				fmt.Printf("%c", ch)
			}
		}
	}
}
