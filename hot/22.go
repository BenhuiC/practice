package hot

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var bk func(s string, l, r int)
	bk = func(s string, l, r int) {
		if l == n && r == n {
			res = append(res, s)
			return
		}
		if l > n {
			return
		}
		if l < n {
			bk(s+"(", l+1, r)
		}
		if r < l {
			bk(s+")", l, r+1)
		}
	}

	bk("", 0, 0)
	return res
}
