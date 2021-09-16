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
		fmt.Println(`usage:
    ./main [flag] < [file]
flags:
    -P to preprocess
    -C to count
    -T to get trigrams
    -I to insert words`)
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
				case ch == 'j' || ch == 'J':
					fallthrough
				case ch == 'h' || ch == 'H':
					fmt.Print("I")
				case ch == 'ñ' || ch == 'Ñ':
					fmt.Print("N")
				case ch == 'k' || ch == 'K':
					fmt.Print("L")
				case ch == 'u' || ch == 'U':
					fallthrough
				case ch == 'w' || ch == 'W':
					fmt.Print("V")
				case ch == 'y' || ch == 'Y':
					fmt.Print("Z")
				case ch == 'á' || ch == 'Á':
					fmt.Print("A")
				case ch == 'é' || ch == 'É':
					fmt.Print("E")
				case ch == 'í' || ch == 'Í':
					fmt.Print("I")
				case ch == 'ó' || ch == 'Ó':
					fmt.Print("O")
				case ch == 'ú' || ch == 'Ú':
					fmt.Print("V")
				case 'a' <= ch && ch <= 'z':
					fmt.Printf("%c", ch-32)
				case ' ' <= ch && ch <= '/' || ':' <= ch && ch <= '?' || ch == '\u00A1' || ch == '\n':
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
	if len(args) == 1 && (args[0] == "-T" || args[0] == "-TC") {
		trigram := make([]rune, 3, 3)
		pmatch := make([]rune, 3, 3)
		tcount := make(map[string]int)
		stdin := make([]rune, 0)
		for {
			if _, err := fmt.Fscanf(os.Stdin, "%c", &ch); err != nil {
				if err == io.EOF {
					break
				}
			} else if ch != '\n' {
				stdin = append(stdin, ch)
			}
		}
		for i := range stdin {
			if i+3 > len(stdin) {
				break
			}
			trigram = stdin[i : i+3]
			strTri := string(trigram)
			if _, ok := tcount[strTri]; !ok {
				if args[0] == "-TC" {
					fmt.Printf("%s ", string(strTri))
				} else {
					fmt.Printf("trigram: %s pos: ", string(strTri))
				}
				tcount[strTri] = 0
				previdx := -1
				for j := 0; j <= len(stdin)-3; j++ {
					if i+1 <= j && j <= i+2 {
						continue
					}
					pmatch = stdin[j : j+3]
					// unoptimized
					if string(pmatch) == strTri {
						tcount[strTri] += 1
						if args[0] == "-T" {
							if previdx == -1 {
								fmt.Printf("%d ", j)
							} else {
								fmt.Printf("(d=%d) %d ", j-previdx, j)
							}
						}
						previdx = j
					}
				}
				previdx = -1
				if args[0] == "-TC" {
					fmt.Printf("%d\n", tcount[strTri])
				} else {
					fmt.Printf("count: %d\n", tcount[strTri])
				}
			}
		}
	}
	if len(args) == 1 && args[0] == "-I" {
		counter := 1
		textlen := 0
		for {
			if _, err := fmt.Fscanf(os.Stdin, "%c", &ch); err != nil {
				if err == io.EOF {
					for i := 0; i < 4-(textlen%4); i++ {
						fmt.Printf("X")
					}
					break
				}
			} else {
				if counter%20 == 0 {
					fmt.Printf("%cAQUI", ch)
					textlen += 4
				} else {
					fmt.Printf("%c", ch)
				}
				counter++
				textlen++
			}
		}
	}
}
