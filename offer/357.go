package offer

func countNumbersWithUniqueDigits(n int) int {
	ary := make([]int, 9)
	ary[0] = 1
	ary[1] = 10
	var f func(n int) int
	f = func(n int) int {
		if ary[n] != 0 {
			return ary[n]
		}
		cur := 9
		for i := 9; i > 10-n; i-- {
			cur *= i
		}
		res := cur + f(n-1)
		ary[n] = res
		return res
	}

	res := f(n)
	return res
}

func countNumbersWithUniqueDigits2(n int) int {
	ary := []int{1, 10, 91, 739, 5275, 32491, 168571, 712891, 2345851}
	return ary[n]
}
