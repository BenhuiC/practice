package partice

import "fmt"

func rob(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	tmp := make([]int, len(nums)+3)
	for i := len(nums) - 1; i >= 0; i-- {
		tmp[i] = max(tmp[i+2], tmp[i+3]) + nums[i]
	}
	fmt.Println(tmp)
	return max(tmp[0], tmp[1])
}
