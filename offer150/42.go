package offer150

import "fmt"

func trap(height []int) int {
	n := len(height)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		e := n - i - 1
		if i == 0 {
			left[i] = height[i]
			right[e] = height[e]
		} else {
			left[i] = max(left[i-1], height[i])
			right[e] = max(right[e+1], height[e])
		}
	}
	fmt.Println(left, right)
	res := 0
	for i := 0; i < n; i++ {
		res += min(left[i], right[i]) - height[i]
	}
	return res
}
