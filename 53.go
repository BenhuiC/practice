package partice

func maxSubArray(nums []int) int {
	var sum, res int
	res = nums[0]
	for _, v := range nums {
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}
		res = Max(sum, res)
	}
	return res
}
