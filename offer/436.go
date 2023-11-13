package offer

import "sort"

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	type pair struct {
		v int
		i int
	}
	start := make([]pair, 0, n)
	end := make([]pair, 0, n)
	for i, v := range intervals {
		start = append(start, pair{v[0], i})
		end = append(end, pair{v[1], i})
	}
	sort.Slice(start, func(i, j int) bool { return start[i].v < start[j].v })
	sort.Slice(end, func(i, j int) bool { return end[i].v < end[j].v })

	res := make([]int, n)
	j := 0
	for _, p := range end {
		for j < n && start[j].v < p.v {
			j++
		}
		if j < n {
			res[p.i] = start[j].i
		} else {
			res[p.i] = -1
		}
	}

	return res
}
