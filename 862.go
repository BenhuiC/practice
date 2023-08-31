package partice

func shortestSubarray(nums []int, k int) int {
	res862 := len(nums) + 1
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = nums[i] + sums[i-1]
	}
	q := make([]int, 0)

	var ed int
	for ; ed < len(nums); ed++ {
		for len(q) != 0 && sums[q[len(q)-1]] >= sums[ed] {
			q = q[:len(q)-1]
		}

		if sums[ed] >= k && res862 > ed+1 {
			res862 = ed + 1
		}
		for len(q) != 0 && sums[ed]-sums[q[0]] >= k {
			if ed-q[0] < res862 {
				res862 = ed - q[0]
			}
			q = q[1:]
		}

		q = append(q, ed)

	}

	if res862 == len(nums)+1 {
		res862 = -1
	}
	return res862
}
