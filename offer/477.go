package offer

func totalHammingDistance(nums []int) int {
	oneCnt := make([]int, 32)
	for _, n := range nums {
		idx := 0
		for n > 0 {
			if n&1 == 1 {
				oneCnt[idx]++
			}
			idx++
			n = n >> 1
		}
	}

	res := 0
	nl := len(nums)
	for _, n := range oneCnt {
		res += n * (nl - n)
	}

	return res
}
