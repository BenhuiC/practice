package offer150

import "fmt"

func longestCommonPrefix(strs []string) string {
	res := ""
	if len(strs) == 0 {
		return res
	}

	idx := 0
	for {
		same := true
		for i, str := range strs {
			if idx >= len(str) {
				same = false
				break
			}
			if i == 0 || str[idx] == strs[i-1][idx] {
				continue
			}
			same = false
			break
		}
		if !same {
			break
		}
		res += fmt.Sprintf("%c", strs[0][idx])
		idx++
	}

	return res
}
