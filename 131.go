package partice

var parResult [][]string

func partition2(s string) [][]string {
	parResult = make([][]string, 0)
	if len(s) == 0 {
		return parResult
	}
	a := make([]string, 0)
	par(s, a, 0)
	return parResult
}

func par(s string, ary []string, index int) {
	if index >= len(s) {
		tmp := make([]string, len(ary))
		copy(tmp, ary)
		parResult = append(parResult, tmp)
		return
	}
	for i := index + 1; i <= len(s); i++ {
		if isHui(s[index:i]) {
			ary = append(ary, s[index:i])
			par(s, ary, i)
			ary = ary[:len(ary)-1]
		}
		continue
	}
}
