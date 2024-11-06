package offer150

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i, j := 0, len(nums)-1
	n := 0
	for i <= j {
		if nums[i] != val {
			i++
		} else if nums[j] == val {
			j--
			n++
		} else {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return len(nums) - n
}
