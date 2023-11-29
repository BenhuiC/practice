package offer

import "sort"

func findLUSlength522(strs []string) int {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) > len(strs[j])
	})
	// s1是s2的子序列
	isSubSeq := func(s1, s2 string) bool {
		if s1 == s2 {
			return true
		}
		i1, i2 := 0, 0
		for i1 < len(s1) && i2 < len(s2) {
			if s1[i1] != s2[i2] {
				i2++
				continue
			}
			i1++
			i2++
		}
		return i1 >= len(s1)
	}
	isSubSeqOfOthers := func(idx int) bool {
		for i := range strs {
			if i == idx {
				continue
			}
			if len(strs[i]) < len(strs[idx]) {
				break
			}
			if isSubSeq(strs[idx], strs[i]) {
				return true
			}
		}
		return false
	}
	for i := range strs {
		if !isSubSeqOfOthers(i) {
			return len(strs[i])
		}
	}

	return -1
}
