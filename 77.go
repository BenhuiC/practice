package partice

var res77 [][]int

func combine(n int, k int) [][]int {
	res77 = make([][]int, 0)
	tryComb(1, n, k, []int{})
	return res77
}

func tryComb(index, n, k int, t []int) {
	if len(t) == k {
		tmp := make([]int, k)
		copy(tmp, t)
		res77 = append(res77, tmp)
		return
	}
	for i := index; i <= n; i++ {
		t = append(t, i)
		tryComb(i+1, n, k, t)
		t = t[:len(t)-1]
	}
}
