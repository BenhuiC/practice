package monotonous_stack

import "fmt"

func smallestSubsequence(s string) string {
	if len(s) == 0 {
		return ""
	}
	chIdx := [26]int{}
	add := [26]bool{}
	for i := range s {
		chIdx[s[i]-'a'] = i
	}

	chs := make([]byte, 0, len(s))
	for i := 0; i < len(s); {
		if len(chs) == 0 {
			add[s[i]-'a'] = true
			chs = append(chs, s[i])
			i++
			continue
		}
		if !add[s[i]-'a'] {
			for len(chs) > 0 && chs[len(chs)-1] > s[i] && chIdx[chs[len(chs)-1]-'a'] > i {
				add[chs[len(chs)-1]-'a'] = false
				chs = chs[:len(chs)-1]
			}
			chs = append(chs, s[i])
			add[s[i]-'a'] = true
		}
		i++

	}

	res := ""
	for _, v := range chs {
		res += fmt.Sprintf("%c", v)
	}

	return res
}
