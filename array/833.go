package array

import "sort"

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	n, m := len(s), len(indices)

	ops := make([][]int, m)
	for i := 0; i < m; i++ {
		ops[i] = []int{i, indices[i]}
	}
	sort.Slice(ops, func(i int, j int) bool {
		return ops[i][1] < ops[j][1]
	})

	res := make([]byte, 0)
	idx := 0
	for i := 0; i < n; {
		if idx >= m || i < ops[idx][1] {
			res = append(res, s[i])
			i++
			continue
		}
		match := false
		for idx < m && i == ops[idx][1] {
			strIdx := ops[idx][0]
			sor := sources[strIdx]
			if i+len(sor) <= n && s[i:i+len(sor)] == sor {
				res = append(res, []byte(targets[strIdx])...)
				i += len(sor)
				match = true
			}
			idx++
		}
		if !match {
			res = append(res, s[i])
			i++
		}
	}

	return string(res)
}
