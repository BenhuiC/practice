package partice

func distributeCandies(candyType []int) int {
	var res575 int
	m := make(map[int]int)
	for _, v := range candyType {
		m[v]++
	}
	if len(m) <= len(candyType)/2 {
		res575 = len(m)
	} else {
		res575 = len(candyType) / 2
	}

	return res575
}
