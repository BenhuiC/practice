package partice

func productExceptSelf(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	left[0] = nums[0]
	right[n-1] = nums[n-1]
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i]
		right[n-i-1] = right[n-i] * nums[n-i-1]
	}
	res238 := make([]int, n)
	for i := range nums {
		if i == 0 {
			res238[i] = right[i+1]
		} else if i == n-1 {
			res238[i] = left[i-1]
		} else {
			res238[i] = left[i-1] * right[i+1]
		}
	}
	return res238
}
