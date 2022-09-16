package partice

import "fmt"

func removeDuplicates80(nums []int) int {
	res80 := len(nums)
	for i := 0; i < len(nums); {
		j := i + 1
		for j < len(nums) && nums[j] == nums[j-1] {
			j++
		}
		// 多出的数量
		d := j - i - 2
		if d <= 0 {
			i = j
			continue
		}
		i = j - d
		res80 -= d
		copy(nums[j-d:], nums[j:])
		nums = nums[:res80]
		fmt.Println(nums)
	}
	return res80
}
