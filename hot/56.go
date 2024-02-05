package hot

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	l, r := 0, 0
	n := len(intervals)
	res := make([][]int, 0)
	for l < n {
		start, end := intervals[l][0], intervals[l][1]
		for r < n && intervals[r][0] <= end {
			if intervals[r][1] > end {
				end = intervals[r][1]
			}
			r++
		}
		res = append(res, []int{start, end})
		l = r
	}
	return res
}
