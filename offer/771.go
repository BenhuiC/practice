package offer

func numJewelsInStones(jewels string, stones string) int {
	m := make(map[byte]bool)
	for _, c := range jewels {
		m[byte(c)] = true
	}
	res := 0
	for _, c := range stones {
		if m[byte(c)] {
			res++
		}
	}

	return res
}
