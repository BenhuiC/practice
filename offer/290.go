package offer

import "strings"

func wordPattern(pattern string, s string) bool {
	sary := strings.Split(s, " ")
	if len(pattern) != len(sary) {
		return false
	}
	pidx, sidx := make(map[byte]int), make(map[string]int)
	for i := range pattern {
		if pidx[pattern[i]] != sidx[sary[i]] {
			return false
		}
		pidx[pattern[i]] = i + 1
		sidx[sary[i]] = i + 1
	}
	return true
}
