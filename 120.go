package partice

import (
	"math"
)

// 动态规划
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return triangle[0][0]
	}
	res120 := math.MaxInt
	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]
			} else if j == i {
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				triangle[i][j] += Min(triangle[i-1][j-1], triangle[i-1][j])
			}
		}
	}
	for i := 0; i < n; i++ {
		res120 = Min(res120, triangle[n-1][i])
	}
	return res120
}
