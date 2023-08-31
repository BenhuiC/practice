package partice

func maxProduct(nums []int) int {
	a, b := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > b {
			a, b = b, nums[i]
			continue
		}
		if nums[i] > a {
			a = nums[i]
		}
	}
	return (a - 1) * (b - 1)
}
