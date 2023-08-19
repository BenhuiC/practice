package partice

func numTimesAllBlue(flips []int) int {
	prefixLen := 0
	res1375 := 0
	for i := 0; i < len(flips); i++ {
		if prefixLen == 0 {
			prefixLen = flips[i]
		} else if prefixLen < flips[i] {
			prefixLen = flips[i]
		}

		if i+1 == prefixLen {
			res1375++
		}
	}
	return res1375
}
