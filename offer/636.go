package offer

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	type pair struct{ fn, t int }
	stack := make([]pair, 0, n)
	res := make([]int, n)
	for _, l := range logs {
		ary := strings.Split(l, ":")
		fn, _ := strconv.Atoi(ary[0])
		idx, _ := strconv.Atoi(ary[2])
		if ary[1] == "start" {
			if len(stack) > 0 {
				res[stack[len(stack)-1].fn] += idx - stack[len(stack)-1].t
			}
			stack = append(stack, pair{fn, idx})
			continue
		}
		sl := len(stack)
		last := stack[sl-1]
		stack = stack[:sl-1]
		res[fn] += idx - last.t + 1
		if len(stack) > 0 {
			stack[sl-2].t = idx + 1
		}
	}

	return res
}
