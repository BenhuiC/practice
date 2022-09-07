package partice

func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	l, r := 1, 2
	sum := 1
	for r <= (target/2)+2 {
		if sum < target {
			sum += r
			r++
		}
		if sum == target {
			t := make([]int, 0, r-l)
			for i := l; i < r; i++ {
				t = append(t, i)
			}
			res = append(res, t)
			sum += r
			r++
		}
		if sum > target {
			sum -= l
			l++
		}
	}

	return res
}
