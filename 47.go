package partice

import "sort"

var puresult [][]int

func permuteUnique(nums []int) [][]int {
	puresult = make([][]int, 0)
	if len(nums) == 0 {
		return puresult
	}
	sort.Ints(nums)
	per(nums, 0)
	return puresult
}

func per(nums []int, index int) {
	if index == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		puresult = append(puresult, tmp)
		return
	}
	for i := index; i < len(nums); i++ {
		if i != index && nums[i-1] == nums[i] {
			continue
		}
		nums[index], nums[i] = nums[i], nums[index]
		per(nums, index+1)
		nums[i], nums[index] = nums[index], nums[i]
	}
}
