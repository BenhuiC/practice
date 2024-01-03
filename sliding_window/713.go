package sliding_window

func numSubarrayProductLessThanK(nums []int, k int) int {
	res := 0
	i, mul := 0, 1
	for j, v := range nums {
		mul *= v
		for i <= j && mul >= k {
			mul /= nums[i]
			i++
		}
		res += j - i + 1
	}

	return res
}
