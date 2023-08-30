package array

func shipWithinDays(weights []int, days int) int {
	var l, r int
	for _, v := range weights {
		if v > l {
			l = v
		}
		r += v
	}

	var can func(w int) bool
	can = func(w int) bool {
		need := 1
		sum := 0
		for _, v := range weights {
			if sum+v > w {
				need++
				sum = 0
			}
			sum += v
		}
		return need <= days
	}

	for l < r {
		mid := l + (r-l)/2
		if can(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
