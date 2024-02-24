package hot

func generate(numRows int) [][]int {
	res := make([][]int, 0, numRows)
	cur := []int{1}
	res = append(res, cur)
	for i := 1; i < numRows; i++ {
		next := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j < len(cur) {
				next[j] += cur[j]
			}
			if j-1 >= 0 {
				next[j] += cur[j-1]
			}
		}
		cur = next
		res = append(res, cur)
	}

	return res
}
