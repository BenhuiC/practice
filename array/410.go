package array

func splitArray(nums []int, k int) int {
	l, r := 0, 0
	for _, v := range nums {
		if v > l {
			l = v
		}
		r += v
	}
	n := len(nums)
	if k >= n {
		return l
	}

	var can func(p int) bool
	can = func(p int) bool {
		need := 1
		sum := 0
		for _, v := range nums {
			if sum+v > p {
				need++
				sum = 0
			}
			sum += v
		}
		return need <= k
	}

	var mid int
	for l < r {
		mid = l + (r-l)/2
		if can(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
