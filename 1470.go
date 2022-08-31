package partice

func shuffle(nums []int, n int) []int {
	tmp := make([]int, n)
	copy(tmp, nums)
	i, j := 0, n
	t := 0
	for i < n && j < 2*n {
		nums[t] = tmp[i]
		nums[t+1] = nums[j]
		t += 2
		i++
		j++
	}

	return nums
}
