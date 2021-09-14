package main

import (
	"fmt"
	"io"
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
			switch ch {
			case 'j':
				fallthrough
			case 'h':
				fmt.Print("i")
			case 'ñ':
				fmt.Print("n")
			case 'k':
				fmt.Print("l")
			case 'u':
				fallthrough
			case 'w':
				fmt.Print("v")
			case 'y':
				fmt.Print("z")
			case 'á':
				fmt.Print("a")
			case 'é':
				fmt.Print("e")
			case 'í':
				fmt.Print("i")
			case 'ó':
				fmt.Print("o")
			case 'ú':
				fmt.Print("v")
			default:
				fmt.Printf("%c", ch)
			}
		}
	}
}
