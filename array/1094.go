package array

func carPooling(trips [][]int, capacity int) bool {
	ary := make([]int, 1001)
	for _, v := range trips {
		val := v[0]
		l, r := v[1], v[2]
		ary[l] += val
		ary[r] -= val
	}
	if ary[0] > capacity {
		return false
	}
	for i := 1; i < len(ary); i++ {
		ary[i] = ary[i] + ary[i-1]
		if ary[i] > capacity {
			return false
		}
	}

	return true
}
