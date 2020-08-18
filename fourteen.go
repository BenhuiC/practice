package partice

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) <= 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	result := ""
	for i := 0; i < len(strs[0]); i++ {
		b := false
		v := strs[0][i]
		for j := 0; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != v {
				b = true
				break
			}
		}
		if b {
			result = strs[0][0:i]
			break
		} else {
			result = strs[0][0 : i+1]
		}
	}
	return result
}
