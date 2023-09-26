package clever

func hammingWeight(num uint32) int {
	if num == 0 {
		return 0
	}
	res := 0
	for num != 0 {
		if num&1 == 1 {
			res++
		}
		num = num >> 1
	}

	return res
}
