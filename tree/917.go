package tree

import "unicode"

func reverseOnlyLetters(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	res := []byte(s)
	l, r := 0, n-1
	for l < r {
		for l < n && !unicode.IsLetter(rune(s[l])) {
			l++
		}
		for r >= 0 && !unicode.IsLetter(rune(s[r])) {
			r--
		}
		if l < r {
			res[l], res[r] = res[r], res[l]
			l++
			r--
		}
	}

	return string(res)
}
