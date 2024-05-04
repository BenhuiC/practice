package array

import "fmt"

func commonChars(words []string) []string {
	common := [26]int{}
	for i := 0; i < 26; i++ {
		common[i] = len(words)
	}
	genChar := func(w string) [26]int {
		r := [26]int{}
		for _, c := range w {
			r[c-'a']++
		}
		return r
	}
	for _, w := range words {
		chars := genChar(w)
		for i := 0; i < 26; i++ {
			common[i] = min(chars[i], common[i])
		}
	}

	res := make([]string, 0)
	for i := 0; i < 26; i++ {
		if common[i] == 0 {
			continue
		}
		for j := 0; j < common[i]; j++ {
			res = append(res, fmt.Sprintf("%c", 'a'+i))
		}
	}

	return res
}
