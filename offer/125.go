package offer

import "strings"

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	s = strings.ToLower(s)
	for l < r {
		if (s[l] < 'a' || s[l] > 'z') && (s[l] < '0' || s[l] > '9') {
			l++
		} else if (s[r] < 'a' || s[r] > 'z') && (s[r] < '0' || s[r] > '9') {
			r--
		} else {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
	}
	return true
}
