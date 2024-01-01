package offer

func selfDividingNumbers(left int, right int) []int {
	isSelfDividing := func(n int) bool {
		for t := n; t != 0; t /= 10 {
			mod := t % 10
			if mod == 0 || n%mod != 0 {
				return false
			}
		}
		return true
	}
	res := make([]int, 0)
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			res = append(res, i)
		}
	}
	return res
}
