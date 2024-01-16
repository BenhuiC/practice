package offer

func letterCasePermutation(s string) []string {
	n := len(s)
	res := make([]string, 0)
	ary := []byte(s)
	var backTrace func(idx int)
	backTrace = func(idx int) {
		if idx == n {
			res = append(res, string(ary))
			return
		}
		backTrace(idx + 1)
		if ary[idx] <= 'z' && ary[idx] >= 'a' {
			ary[idx] = ary[idx] - 32
			backTrace(idx + 1)
		} else if ary[idx] <= 'Z' && ary[idx] >= 'A' {
			ary[idx] = ary[idx] + 32
			backTrace(idx + 1)
		}
	}
	backTrace(0)
	return res
}
