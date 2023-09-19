package back_trace

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	var backTrace func(ary []int, idx int)
	backTrace = func(ary []int, idx int) {
		if idx == len(nums) {
			r := make([]int, len(nums))
			copy(r, ary)
			res = append(res, r)
			return
		}
		for i := idx; i < len(nums); i++ {
			ary[idx], ary[i] = ary[i], ary[idx]
			backTrace(ary, idx+1)
			ary[idx], ary[i] = ary[i], ary[idx]
		}
	}
	backTrace(nums, 0)
	return res
}
