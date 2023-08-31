package partice

func minDeletionSize955(strs []string) int {
	n := len(strs)
	if n <= 1 {
		return 0
	}
	res := 0
	strLen := len(strs[0])
	cmp := make(map[int]struct{})

	for idx := 0; idx < strLen; idx++ {
		del := false
		nextCmp := make(map[int]struct{})
		for j := 1; j < n; j++ {
			if _, ok := cmp[j]; ok {
				nextCmp[j] = struct{}{}
				continue
			}
			if strs[j-1][idx] < strs[j][idx] {
				nextCmp[j] = struct{}{}
			} else if strs[j-1][idx] > strs[j][idx] {
				del = true
			}
		}
		if del {
			res++
		} else {
			cmp = nextCmp
		}
	}

	return res
}
