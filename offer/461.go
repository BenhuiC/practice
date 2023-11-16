package offer

func hammingDistance(x int, y int) int {
	res := 0
	x = x ^ y
	for x != 0 {
		res++
		x = x & (x - 1)
	}

	return res
}
