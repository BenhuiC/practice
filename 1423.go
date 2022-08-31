package partice

func maxScore(cardPoints []int, k int) int {
	if k == 1 {
		return max(cardPoints[0], cardPoints[len(cardPoints)-1])
	}
	res := 0
	for i := 0; i < k; i++ {
		res += cardPoints[i]
	}
	if k == len(cardPoints) {
		return res
	}
	l := len(cardPoints) - 1
	cur := res // 窗口内和
	for r := k - 1; r >= 0; r-- {
		cur -= cardPoints[r]
		cur += cardPoints[l]
		res = max(cur, res)
		l--
	}

	return res
}
