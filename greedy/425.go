package greedy

import "sort"

func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n <= 1 {
		return 1
	}
	sort.Slice(points, func(i, j int) bool {
		//if points[i][0] != points[j][0] {
		//	return points[i][0] < points[j][0]
		//}
		return points[i][1] < points[j][1]
	})
	res := 0
	idx := 0
	for idx < n {
		j := idx + 1
		for j < n && points[j][0] <= points[idx][1] {
			j++
		}
		idx = j
		res++
	}

	return res
}
