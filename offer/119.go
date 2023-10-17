package offer

func getRow(rowIndex int) []int {
	res := make([]int, 0, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		ary := make([]int, i+1)
		ary[0], ary[i] = 1, 1
		for x := 1; x < i; x++ {
			ary[x] = res[x] + res[x-1]
		}
		res = ary
	}

	return res
}
