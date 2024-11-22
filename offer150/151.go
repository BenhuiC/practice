package offer150

import "strings"

func reverseWords(s string) string {
	ary := strings.Split(strings.TrimSpace(s), " ")
	res := make([]string, 0, len(ary))
	for i := len(ary) - 1; i >= 0; i-- {
		if ary[i] == "" {
			continue
		}
		res = append(res, ary[i])
	}
	return strings.Join(res, " ")
}
