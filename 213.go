package partice

import "fmt"

func rob2(nums []int) int {
	res213 := 0
	if len(nums) <= 3 {
		for _, v := range nums {
			if v > res213 {
				res213 = v
			}
		}
		return res213
	}
	return max(robRg(nums[1:]), robRg(nums[:len(nums)-1]))
}

func robRg(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	tmp := make([]int, len(nums)+2)
	for i := len(nums) - 1; i >= 0; i-- {
		tmp[i] = max(tmp[i+1], tmp[i+2]+nums[i])
	}
	fmt.Println(tmp)
	return max(tmp[0], tmp[1])
}
