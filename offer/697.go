package offer

func findShortestSubArray(nums []int) int {
	res := len(nums)
	mp := make(map[int]*[3]int)
	du := 0
	for i, v := range nums {
		m, ok := mp[v]
		if !ok {
			m = &[3]int{i, i, 1}
			mp[v] = m
		} else {
			m[1] = i
			m[2]++
		}
		if m[2] > du {
			du = m[2]
			res = m[1] - m[0] + 1
		} else if tmp := m[1] - m[0] + 1; m[2] == du && tmp < res {
			res = tmp
		}
	}

	return res
}
