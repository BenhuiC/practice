package greedy

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n < 2 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	ln, right := 1, intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][0] >= right {
			ln++
			right = intervals[i][1]
		}
	}

	return n - ln
}
