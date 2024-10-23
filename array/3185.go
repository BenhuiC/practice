package array

func countCompleteDayPairs3185(hours []int) int64 {
	var res int64
	m := make(map[int]int64)
	for _, v := range hours {
		t := v % 24
		res += m[(24-t)%24]
		m[t]++
	}

	return res
}
