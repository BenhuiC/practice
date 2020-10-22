package partice

import "sort"

var threeResult [][]int

func threeSum(nums []int) [][]int {
	threeResult = make([][]int, 0)
	if len(nums) < 3 {
		return threeResult
	}
	sort.Ints(nums)
	three(nums, make([]int, 0), 0, 0)
	return threeResult
}

func three(nums, ary []int, index, sum int) {
	if len(ary) == 2 {
		for i := index; i < len(nums); i++ {
			if sum+nums[i] == 0 {
				tmp := make([]int, 0)
				tmp = append(tmp, ary...)
				tmp = append(tmp, nums[i])
				threeResult = append(threeResult, tmp)
				return
			}
		}
	} else {
		for i := index; i < len(nums)-1; i++ {
			if i != index && nums[i] == nums[i-1] {
				continue
			}
			sum += nums[i]
			ary = append(ary, nums[i])
			three(nums, ary, i+1, sum)
			ary = ary[0 : len(ary)-1]
			sum -= nums[i]
		}
	}
}
