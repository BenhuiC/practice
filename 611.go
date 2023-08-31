package partice

import "sort"

func triangleNumber(nums []int) int {
	res611 := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		k := i + 1
		for j := i + 1; j < len(nums); j++ {
			for k+1 < len(nums) && nums[k+1] < nums[i]+nums[j] {
				k++
			}
			res611 += Max(k-j, 0)
		}
	}
	return res611
}
