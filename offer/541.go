package offer

func reverseStr(s string, k int) string {
	if k == 1 {
		return s
	}
	ary := []byte(s)
	n := len(s)
	reverse := func(i, j int) {
		for i <= j {
			ary[i], ary[j] = ary[j], ary[i]
			i++
			j--
		}
	}
	var l, r int
	for l < len(s) {
		r = l + k - 1
		if r >= n {
			r = n - 1
		}
		reverse(l, r)

		l += 2 * k
	}
	return string(ary)
}
