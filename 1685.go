package partice

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	t := nums[0]
	nums[0] = 0
	for i := 1; i < n; i++ {
		nums[0] += nums[i] - t
	}

	for i := 1; i < n; i++ {
		sub := nums[i] - t
		t = nums[i]
		//nums[i] = nums[i-1] - (n-i-1)*sub + (i-1)*sub
		nums[i] = nums[i-1] + (2*i-n)*sub
	}

	return nums
}

/*
i-1-n+i+1 =2*i-n
*/
