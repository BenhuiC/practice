package offer

func magicalString(n int) int {
	if n < 4 {
		return 1
	}
	ary := make([]byte, n)
	copy(ary, "122")

	i, j := 2, 3
	res := 1
	for j < n {
		size := ary[i] - '0'
		num := '3' - ary[j-1]
		for size > 0 && j < n {
			ary[j] = '0' + num
			j++
			size--
			if num == 1 {
				res++
			}
		}
		i++
	}

	return res
}
