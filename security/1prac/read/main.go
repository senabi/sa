package main

import (
	"bufio"
	"fmt"
	"io"
	//"log"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	//defer writer.Flush()
	var ch rune
	for {
		if _, err := fmt.Fscanf(os.Stdin, "%c", &ch); err != nil {
			//if ch, _, err := reader.ReadRune(); err != nil {
			//if ch, err := reader.ReadByte(); err != nil {
			if err == io.EOF {
				break
			}
		} else {
			//writer.WriteByte(ch)
			//fmt.Printf("%q, %U, %d, %d\n", ch, ch, ch, sz)
			//writer.WriteRune(ch)
			//fmt.Fprintf(writer, "%d %c\n", ch, ch)
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
			case '\u00E1':
				fmt.Print("a")
			case 'é':
				fmt.Print("e")
			case 'í':
				fmt.Print("i")
			case 'ó':
				fmt.Print("o")
			case 'ú':
				fmt.Print("u")
			default:
				fmt.Printf("%c", ch)
			}
			//fmt.Fprintf(os.Stdout, "%d %c\n", ch, ch)
			//fmt.Printf("%q, %d, %08b\n", ch, ch, ch)
		}
	}
}
