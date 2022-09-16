package partice

import "fmt"

func numDecodings(s string) int {
	n := len(s)
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			f[i] += f[i-1]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			f[i] += f[i-2]
		}
	}
	return f[n]
}

func numDecodings2(s string) int {
	res91 := 0
	m := map[string]struct{}{}
	for i := 1; i <= 26; i++ {
		m[fmt.Sprint(i)] = struct{}{}
	}
	var backt func(v int)
	backt = func(v int) {
		if v >= len(s) {
			res91++
			return
		}
		for i := v; i <= v+1 && i < len(s); i++ {
			if _, ok := m[s[v:i+1]]; ok {
				backt(i + 1)
			} else {
				break
			}
		}
	}
	backt(0)
	return res91
}
