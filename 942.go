package partice

/*
示例 1：

输入：s = "IDID"
输出：[0,4,1,3,2]
示例 2：

输入：s = "III"
输出：[0,1,2,3]
示例 3：

输入：s = "DDI"
输出：[3,2,0,1]
*/

func diStringMatch(s string) []int {
	res := make([]int, 0, len(s))
	if s == "" {
		return res
	}
	l, r := 0, len(s)
	for _, v := range s {
		tmp := 0
		if v == 'I' {
			tmp = l
			l++
		} else {
			tmp = r
			r--
		}
		res = append(res, tmp)
	}
	if s[len(s)-1] == 'I' {
		res = append(res, r)
	} else {
		res = append(res, l)
	}

	return res
}
