package array

func sortArrayByParityII(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	odd, even := 0, 1
	for {
		for odd < n && nums[odd]&1 == 0 {
			odd += 2
		}
		for even < n && nums[even]&1 == 1 {
			even += 2
		}
		if odd < n && even < n {
			nums[odd], nums[even] = nums[even], nums[odd]
		} else {
			break
		}
	}
	return nums
}
