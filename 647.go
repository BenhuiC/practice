package partice

func countSubstrings(s string) int {
	res647 := 0
	n := len(s)
	for i := 0; i < n; i++ {
		res647++
		l, h := i-1, i+1
		for l >= 0 && h < n && s[l] == s[h] {
			res647++
			l--
			h++
		}
		if i != n-1 && s[i] == s[i+1] {
			res647++
			l, h = i-1, i+2
			for l >= 0 && h < n && s[l] == s[h] {
				res647++
				l--
				h++
			}
		}
	}

	return res647
}
