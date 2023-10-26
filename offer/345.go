package offer

func reverseVowels(s string) string {
	ary := []byte(s)
	m := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !m[ary[l]] {
			l++
		}
		for r > l && !m[ary[r]] {
			r--
		}
		if l < r {
			ary[l], ary[r] = ary[r], ary[l]
			l++
			r--
		}
	}

	return string(ary)
}
