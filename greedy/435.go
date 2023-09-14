package greedy

import "sort"

// not pass
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n < 2 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] < intervals[j][1]
		}
	})

	res := 0
	l, r := 0, 1
	for l < r && r < n {
		if intervals[r][0] < intervals[l][1] {
			r++
			res++
		} else {
			l = r
			r++
		}
	}

	return res
}
