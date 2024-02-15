package hot

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	var bk func(ary []int, idx int)
	bk = func(ary []int, idx int) {
		res = append(res, append([]int{}, ary...))
		for i := idx; i < len(nums); i++ {
			ary = append(ary, nums[i])
			bk(ary, i+1)
			ary = ary[:len(ary)-1]
		}

	}
	bk([]int{}, 0)

	return res
}
