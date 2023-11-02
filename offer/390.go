package offer

func lastRemaining(n int) int {
	var solution func(n int, left bool) int
	solution = func(n int, left bool) int {
		if n == 1 {
			return 1
		} else if n <= 5 && left {
			return 2
		}
		if left || n%2 == 1 {
			return 2 * solution(n/2, !left)
		}
		return 2*solution(n/2, !left) - 1
	}
	return solution(n, true)
}

func lastRemaining2(n int) int {
	head := 1
	steps := 1
	left := true
	for n > 1 {
		// 从左边开始移除 or（从右边开始移除，数列总数为奇数）
		if left || n%2 != 0 {
			head += steps
		}

		steps <<= 1  //步长 * 2
		n >>= 1      //总数 / 2
		left = !left //取反移除方向
	}

	return head
}
