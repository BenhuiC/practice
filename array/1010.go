package array

func numPairsDivisibleBy60(time []int) int {
	prefix := make(map[int]int)
	res := 0
	for _, v := range time {
		y := v % 60
		if y == 0 {
			res += prefix[0]
		} else {
			res += prefix[60-y]
		}

		prefix[y]++
	}

	return res
}
