package array

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	res := 1
	idx := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[idx] = nums[i]
			idx++
			res++
		}
	}
	return res
}
