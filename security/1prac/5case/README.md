```go
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
```
