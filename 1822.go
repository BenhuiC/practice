package partice

func arraySign(nums []int) int {
	res := -1
	pos := true
	for _, v := range nums {
		if v == 0 {
			return 0
		} else if v < 0 {
			pos = !pos
		}
	}
	if pos {
		res = 1
	}
	return res
}
