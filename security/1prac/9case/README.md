```go
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
```
