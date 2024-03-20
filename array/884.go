package array

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	m1 := make(map[string]int)
	for _, s := range strings.Split(s1, " ") {
		m1[s]++
	}
	for _, s := range strings.Split(s2, " ") {
		m1[s]++
	}

	res := make([]string, 0)
	for k, v := range m1 {
		if v > 1 {
			continue
		}
		res = append(res, k)
	}
	return res
}
