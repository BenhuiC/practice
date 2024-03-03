package hot

func singleNumber(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	v := nums[0]

	for i := 1; i < n; i++ {
		v = v ^ nums[i]
	}

	return v
}
