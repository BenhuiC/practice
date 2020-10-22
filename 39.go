package partice

var combResult [][]int

func combinationSum(candidates []int, target int) [][]int {
	combResult = make([][]int, 0)
	a := make([]int, 0)
	comb(candidates, a, 0, 0, target)
	return combResult
}

func comb(cand, ary []int, index, sum, target int) {
	if sum == target {
		tmp := make([]int, len(ary))
		copy(tmp, ary)
		combResult = append(combResult, tmp)
	} else if sum > target {
		return
	}
	for i := index; i < len(cand); i++ {
		sum += cand[i]
		ary = append(ary, cand[i])
		comb(cand, ary, i, sum, target)
		sum -= cand[i]
		ary = ary[:len(ary)-1]
	}
}
