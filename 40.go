package partice

import "sort"

var com2Result [][]int

func combinationSum2(candidates []int, target int) [][]int {
	com2Result = make([][]int, 0)
	if len(candidates) == 0 {
		return com2Result
	}
	sort.Ints(candidates)
	com2(candidates, []int{}, 0, 0, target)
	return com2Result
}

func com2(candidates, ap []int, index, sum, target int) {
	if sum > target {
		return
	}
	if target == sum {
		tmp := make([]int, len(ap))
		copy(tmp, ap)
		com2Result = append(com2Result, tmp)
		return
	}
	for i := index; i < len(candidates); i++ {
		if i != index && candidates[i] == candidates[i-1] {
			continue
		}
		ap = append(ap, candidates[i])
		com2(candidates, ap, i+1, sum+candidates[i], target)
		ap = ap[:len(ap)-1]
	}
}
