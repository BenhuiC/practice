package partice

func numsSameConsecDiff(n int, k int) []int {
	res := make([]int, 0)
	for i := 1; i < 10; i++ {
		if i+k < 10 || i-k >= 0 {
			res = append(res, i)
		}
	}
	start := pow10(n - 1)
	for {
		if res[0] >= start {
			break
		}
		tmpVal := res[0]
		lastDig := res[0] % 10
		res = res[1:]
		if lastDig-k >= 0 {
			res = append(res, tmpVal*10+lastDig-k)
		}
		if k == 0 {
			continue
		}
		if lastDig+k < 10 {
			res = append(res, tmpVal*10+lastDig+k)
		}
	}

	return res
}

func pow10(n int) int {
	v := 1
	for i := 0; i < n; i++ {
		v *= 10
	}
	return v
}
