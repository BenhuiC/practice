package partice

import "strings"

func lengthOfLastWord(s string) int {
	result := 0
	s = strings.TrimRight(s, " ")
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		result++
	}
	return result
}
