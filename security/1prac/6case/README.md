```go
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
```
