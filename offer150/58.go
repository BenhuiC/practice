package offer150

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")
	n := len(s)
	spaceIdx := n - 1
	for spaceIdx >= 0 {
		if s[spaceIdx] == ' ' {
			break
		}
		spaceIdx--
	}
	return n - spaceIdx - 1
}
