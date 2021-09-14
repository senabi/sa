package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("usage:\n    -P to preprocess\n    -C to count")
		return
	}
	var ch rune
	if len(args) == 1 && args[0] == "-P" {
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
				case ' ' <= ch && ch <= '/' || ':' <= ch && ch <= '?' || ch == '\u00A1':
					fmt.Print("")
				default:
					fmt.Printf("%c", ch)
				}
			}
		}
	}
	if len(args) == 1 && args[0] == "-C" {
		alphabet := make(map[string]int)
		for {
			if _, err := fmt.Fscanf(os.Stdin, "%c", &ch); err != nil {
				if err == io.EOF {
					break
				}
			} else {
				chs := string(ch)
				if _, ok := alphabet[chs]; ok {
					alphabet[chs] += 1
				} else {
					alphabet[chs] = 0
				}
			}
		}
		type pair struct {
			k string
			v int
		}
		freq := make([]pair, 0, len(alphabet))
		for k, v := range alphabet {
			freq = append(freq, pair{k, v})
		}
		sort.Slice(freq, func(i, j int) bool { return freq[i].v > freq[j].v })
		//5 most freq
		for _, kv := range freq[:5] {
			fmt.Printf("%q : %d\n", kv.k, kv.v)
		}
	}
}
