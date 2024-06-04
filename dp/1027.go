package dp

func longestArithSeqLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	minVal, maxVal := nums[0], nums[0]
	for _, v := range nums {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	diff := maxVal - minVal
	res := 1
	for d := -diff; d <= diff; d++ {
		m := make(map[int]int)
		for _, v := range nums {
			if nu, ok := m[v-d]; ok {
				m[v] = nu + 1
			} else {
				m[v] = 1
			}
			res = max(res, m[v])
		}
	}

	return res
}
