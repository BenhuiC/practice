package array

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}
