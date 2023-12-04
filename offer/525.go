package offer

func findMaxLength(nums []int) int {
	zeroCnt, oneCnt := 0, 0
	res := 0
	m := map[int]int{0: -1}
	for i, v := range nums {
		if v == 0 {
			zeroCnt++
		} else {
			oneCnt++
		}
		diff := zeroCnt - oneCnt
		if idx, has := m[diff]; has {
			if i-idx > res {
				res = i - idx
			}
		} else {
			m[diff] = i
		}
	}
	return res
}
