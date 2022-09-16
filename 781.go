package partice

func numRabbits(answers []int) int {
	var res781 int
	m := make(map[int]int)
	for _, v := range answers {
		m[v]++
	}
	for k, v := range m {
		if v == 1 {
			res781 += k + 1
			continue
		}
		c := v / (k + 1)
		yu := v % (k + 1)
		res781 += c * (k + 1)
		if yu != 0 {
			res781 += k + 1
		}
	}
	return res781
}
