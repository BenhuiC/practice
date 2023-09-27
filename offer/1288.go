package offer

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] > intervals[j][1]
	})
	nlen := len(intervals)
	res := nlen
	cur, next := 0, 0
	for cur < nlen {
		for next = cur + 1; next < nlen; next++ {
			if intervals[next][1] > intervals[cur][1] {
				break
			}
			res--
		}
		cur = next
	}

	return res
}
