package partice

func constructArray(n int, k int) []int {
	res := make([]int, 0, n)
	res = append(res, 1)
	b := 1
	diff := k
	for diff >= 1 {
		res = append(res, res[len(res)-1]+(diff*b))
		diff--
		b *= -1
	}
	for i := k + 2; i <= n; i++ {
		res = append(res, i)
	}
	return res
}

/*
10 5
1 6 2 5 3 4 7 8 9 10

3 2
1 3 2

5 2
1 3 2 4 5
*/
