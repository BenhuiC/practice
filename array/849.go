package array

func maxDistToClosest(seats []int) int {
	n := len(seats)
	if n < 1 {
		return 0
	}
	res := 0
	l, r := 0, 0

	for l < n && seats[l] == 0 {
		l++
	}
	res = l
	for l < n {
		// l  最左侧为1的位置
		for l < n-1 && seats[l+1] == 1 {
			l++
		}

		// r 最右侧为1的位置
		r = l + 1
		for r < n && seats[r] == 0 {
			r++
		}

		diff := r - l
		if r < n {
			diff /= 2
		} else {
			diff--
		}
		res = max(res, diff)

		l = r
	}

	return res
}
