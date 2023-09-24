package back_trace

import "strings"

func generateParenthesis(n int) []string {
	l, r := n, n
	res := make([]string, 0)
	var backtrace func(s string, l, r int)
	backtrace = func(s string, l, r int) {
		if l == 0 {
			s += strings.Repeat(")", r)
			res = append(res, s)
			return
		}
		backtrace(s+"(", l-1, r)
		if r > l {
			backtrace(s+")", l, r-1)
		}
	}
	backtrace("", l, r)

	return res
}
