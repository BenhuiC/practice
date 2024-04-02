package array

func minDeletionSize(strs []string) int {
	m := len(strs)
	if m < 2 {
		return 0
	}
	n := len(strs[0])

	res := 0
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			if strs[j][i] < strs[j-1][i] {
				res++
				break
			}
		}
	}

	return res
}
