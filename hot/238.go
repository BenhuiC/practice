package hot

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	t := 1
	for i, v := range nums {
		t *= v
		res[i] = t
	}

	t = 1
	for j := n - 1; j >= 0; j-- {
		if j == 0 {
			res[j] = t
		} else {
			res[j] = res[j-1] * t
			t *= nums[j]
		}
	}

	return res
}
