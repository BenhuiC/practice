package offer

import (
	"math"
	"sort"
)

func minMoves2(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	res := math.MaxInt
	n := len(nums)
	leftSum := 0
	var l, r int
	for i, v := range nums {
		leftSum += v
		l = i + 1
		r = n - l
		tmp := (l-r)*v + sum - 2*leftSum
		if tmp < res {
			res = tmp
		}
	}

	return res
}

// 直接用中位数
func minMoves2_(nums []int) (ans int) {
	sort.Ints(nums)
	x := nums[len(nums)/2]
	for _, num := range nums {
		ans += abs(num - x)
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
