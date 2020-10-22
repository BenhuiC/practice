package partice

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	var result int
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	result = nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			tmp := nums[i] + nums[j] + nums[k]
			if math.Abs(float64(tmp-target)) < math.Abs(float64(target-result)) {
				result = tmp
			}
			if tmp > target {
				k--
			} else {
				j++
			}
		}
	}

	return result
}
