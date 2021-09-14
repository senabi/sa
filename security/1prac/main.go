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
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("usage:\n    -E to encode\n    -D to decode")
		return
	}
	var ch rune
	if len(args) == 1 && args[0] == "-E" {
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
				case ' ' <= ch && ch <= '/' || ':' <= ch && ch <= '?' || ch == '\u00A1':
					fmt.Print("")
				default:
					fmt.Printf("%c", ch)
				}
				//fmt.Fprintf(os.Stdout, "%d %c\n", ch, ch)
				//fmt.Printf("%q, %d, %08b\n", ch, ch, ch)
			}
		}
	}
}
