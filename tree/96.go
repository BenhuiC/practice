package tree

func numTrees(n int) int {
	if n <= 1 {
		return n
	}
	ary := make([]int, n+1)
	ary[0] = 1
	ary[1] = 1
	ary[2] = 2

	var f func(i int) int
	f = func(i int) int {
		if i <= 0 {
			return 1
		}
		if ary[i] > 0 {
			return ary[i]
		}
		r := 0
		for t := 1; t <= i; t++ {
			r += f(t-1) * f(i-t)
		}
		ary[i] = r
		return r
	}
	return f(n)
}
