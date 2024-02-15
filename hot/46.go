package hot

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)

	used := make([]bool, n)
	var bk func(ary []int)
	bk = func(ary []int) {
		if len(ary) == n {
			dst := make([]int, n)
			copy(dst, ary)
			res = append(res, dst)
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			ary = append(ary, nums[i])
			bk(ary)
			used[i] = false
			ary = ary[:len(ary)-1]
		}
	}

	bk(make([]int, 0, n))
	return res
}
