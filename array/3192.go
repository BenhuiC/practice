package array

func minOperations3192(nums []int) int {
	res := 0
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] != nums[i+1] {
			res++
		}
	}
	if nums[0] == 0 {
		res++
	}

	return res
}
