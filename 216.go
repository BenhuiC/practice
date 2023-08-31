package partice

var res216 [][]int

func combinationSum3(k int, n int) [][]int {
	res216 = make([][]int, 0)
	tryCombSum3(k, n, 1, 0, []int{})
	return res216
}

func tryCombSum3(k, n, nIndex, sum int, t []int) {
	if sum == n && len(t) == k {
		tmp := make([]int, k)
		copy(tmp, t)
		res216 = append(res216, tmp)
	}
	for i := nIndex; i <= 9 && i <= n-sum; i++ {
		t = append(t, i)
		tryCombSum3(k, n, i+1, sum+i, t)
		t = t[:len(t)-1]
	}
}
