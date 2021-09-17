```go
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
```
