package offer150

func removeDuplicates(nums []int) int {
	left, right := 0, 0
	for right < len(nums) {
		if nums[left] == nums[right] {
			right++
		} else {
			nums[left+1] = nums[right]
			left++
		}
	}
	return left + 1
}
