package hot

func partition(s string) [][]string {
	res := make([][]string, 0)

	isPartition := func(str string) bool {
		if str == "" {
			return false
		}
		n := len(str)
		if n == 1 {
			return true
		}
		for l, r := 0, n-1; l <= r; {
			if str[l] != str[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	var bk func(ary []string, idx int)
	bk = func(ary []string, idx int) {
		if idx >= len(s) {
			res = append(res, append([]string{}, ary...))
			return
		}
		for i := idx + 1; i <= len(s); i++ {
			if !isPartition(s[idx:i]) {
				continue
			}
			ary = append(ary, s[idx:i])
			bk(ary, i)
			ary = ary[:len(ary)-1]
		}
	}

	bk([]string{}, 0)
	return res
}
