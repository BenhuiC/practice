package offer

func reverseWords(s string) string {
	ary := []byte(s)
	i := 0
	for i < len(s) {
		if s[i] == ' ' {
			i++
			continue
		}
		l, r := i, i
		for r+1 < len(s) && s[r+1] != ' ' {
			r++
		}
		i = r + 1
		for l < r {
			ary[l], ary[r] = ary[r], ary[l]
			l++
			r--
		}
	}
	return string(ary)
}
