package array

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	n := len(logs)
	if n == 0 {
		return nil
	}
	isNum := func(s string) bool {
		n := len(s)
		if n == 0 {
			return false
		}
		return s[n-1] >= '0' && s[n-1] <= '9'
	}

	l, r := n-1, n-1
	for l >= 0 && r >= 0 {
		for r >= 0 && isNum(logs[r]) {
			r--
		}
		for l >= 0 && !isNum(logs[l]) {
			l--
		}
		if l >= 0 && l < r {
			logs[l], logs[r] = logs[r], logs[l]
		} else if l > r {
			l = r - 1
		}
	}
	sort.Slice(logs[:r+1], func(i, j int) bool {
		l1 := strings.IndexByte(logs[i], ' ')
		l2 := strings.IndexByte(logs[j], ' ')
		s1 := logs[i][l1+1:]
		s2 := logs[j][l2+1:]
		if s1 == s2 {
			return logs[i][:l1] < logs[j][:l2]
		} else {
			return s1 < s2
		}
	})

	return logs
}
