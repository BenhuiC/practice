package partice

func distributeCandies(candyType []int) int {
	m := map[int]struct{}{}
	for _, v := range candyType {
		m[v] = struct{}{}
	}
	res := len(candyType) / 2
	if len(m) < res {
		res = len(m)
	}
	return res
}
