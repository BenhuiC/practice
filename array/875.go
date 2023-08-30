package array

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 0
	for _, v := range piles {
		if v > r {
			r = v
		}
	}
	if h == len(piles) {
		return r
	}
	var can func(d int) bool
	can = func(d int) bool {
		need := 0
		for _, p := range piles {
			need += p / d
			if p%d != 0 {
				need++
			}
		}
		return need <= h
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
