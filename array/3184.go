package array

func countCompleteDayPairs(hours []int) int {
	res := 0
	m := make(map[int]int)
	for _, v := range hours {
		t := v % 24
		res += m[(24-t)%24]
		m[t]++
	}

	return res
}
