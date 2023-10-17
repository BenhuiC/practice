package offer

func generate(numRows int) [][]int {
	res := make([][]int, 0, numRows)
	for i := 0; i < numRows; i++ {
		ary := make([]int, i+1)
		ary[0], ary[i] = 1, 1
		for x := 1; x < i; x++ {
			ary[x] = res[i-1][x] + res[i-1][x-1]
		}
		res = append(res, ary)
	}

	return res
}
