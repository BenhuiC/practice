package clever

func binaryGap(n int) int {
	res := 0
	cur := 0
	for n > 0 {
		if n&1 == 1 {
			res = max(cur, res)
			cur = 1
		} else if cur != 0 {
			cur++
		}
		n = n >> 1
	}
	return res
}
