package array

func beautifulArray(n int) []int {
	res := make([]int, 0, n)
	if n < 3 {
		for i := 1; i <= n; i++ {
			res = append(res, i)
		}
		return res
	}
	left := beautifulArray((n + 1) / 2)
	right := beautifulArray(n / 2)

	for _, v := range left {
		res = append(res, int((float64(v)-0.5)*2))
	}
	for _, v := range right {
		res = append(res, v*2)
	}
	return res
}
