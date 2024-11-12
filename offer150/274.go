package offer150

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)
	n := len(citations)
	res := 0
	for i := n - 1; i >= 0; i-- {
		v := min(citations[i], n-i)
		res = max(res, v)
	}
	return res
}
