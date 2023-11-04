package offer

import "math"

func thirdMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	first, second, third := math.MinInt, math.MinInt, math.MinInt
	for _, v := range nums {
		if v <= third {
			continue
		}
		if v > first {
			first, second, third = v, first, second
		} else if first > v && v > second {
			second, third = v, second
		} else if second > v && v > third {
			third = v
		}
	}
	if third == math.MinInt {
		return first
	}
	return third
}
