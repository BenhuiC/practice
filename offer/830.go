package offer

func largeGroupPositions(s string) [][]int {
	l, r := 0, 0
	res := make([][]int, 0)
	n := len(s)
	for r < n {
		for r+1 < n && s[r+1] == s[r] {
			r++
		}
		if r-l+1 >= 3 {
			res = append(res, []int{l, r})
		}
		l, r = r+1, r+1
	}

	return res
}
