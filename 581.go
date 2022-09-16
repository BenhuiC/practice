package partice

import (
	"math"
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	var res581 int
	n := len(nums)
	ary := append([]int{}, nums...)
	sort.Ints(ary)
	var l, r int
	for l = 0; l < n && ary[l] == nums[l]; l++ {
	}
	if l == n {
		return 0
	}
	for r = n - 1; r >= 0 && ary[r] == nums[r]; r-- {
	}
	res581 = r - l + 1
	return res581
}

func findUnsortedSubarray2(nums []int) int {
	n := len(nums)
	minn, maxn := math.MaxInt64, math.MinInt64
	left, right := -1, -1
	for i, num := range nums {
		// 从左到右便利，maxn记录的是便利过的数组中最大值，如果说当前值是小与最大值，那么他肯定在待排序的数组中
		if maxn > num {
			right = i
		} else {
			maxn = num
		}
		if minn < nums[n-i-1] {
			left = n - i - 1
		} else {
			minn = nums[n-i-1]
		}
	}
	if right == -1 {
		return 0
	}
	return right - left + 1
}
