package offer

import "sort"

func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	l, r := 0, 0
	nlen := len(intervals)
	for l < nlen {
		left := intervals[l][0]
		right := intervals[l][1]
		for r < nlen && intervals[r][0] <= right {
			if intervals[r][1] > right {
				right = intervals[r][1]
			}
			r++
		}
		res = append(res, []int{left, right})
		l = r
	}

	return res
}
