package offer

func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	res := 0
	ary := make([]string, 0)
	ary = append(ary, startGene)
	trace := make(map[string]struct{})
	for len(ary) != 0 {
		res++
		next := make([]string, 0)
		for _, v := range ary {
			for _, b := range bank {
				if v == b {
					continue
				}
				if _, ok := trace[b]; ok {
					continue
				}
				diffNum := 0
				for i := 0; i < 8; i++ {
					if v[i] != b[i] {
						diffNum++
					}
					if diffNum > 1 {
						break
					}
				}
				if diffNum > 1 {
					continue
				}
				if b == endGene {
					return res
				}
				trace[b] = struct{}{}
				next = append(next, b)
			}
		}

		ary = next
	}

	return -1
}
