package offer150

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	mult := 1
	for i, v := range nums {
		mult *= v
		res[i] = mult
	}

	mult = 1
	for i := n - 1; i >= 0; i-- {
		if i > 0 {
			res[i] = mult * res[i-1]
		} else {
			res[i] = mult
		}
		mult *= nums[i]
	}

	return res
}
