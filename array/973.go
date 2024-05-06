package array

import "sort"

func kClosest(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		x1, y1 := points[i][0], points[i][1]
		x2, y2 := points[j][0], points[j][1]
		return x1*x1+y1*y1 < x2*x2+y2*y2
	})
	return points[:k]
}
