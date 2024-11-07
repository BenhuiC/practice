package offer150

func removeDuplicates80(nums []int) int {
	idx, left, right := 0, 0, 0
	for left < len(nums) {
		for right = left + 1; right < len(nums) && nums[left] == nums[right]; right++ {
		}
		nums[idx] = nums[left]
		idx++
		if right-left >= 2 {
			nums[idx] = nums[left]
			idx++
		}
		left = right
	}
	return idx
}
