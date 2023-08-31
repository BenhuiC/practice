package partice

func partitionLabels(s string) []int {
	if len(s) == 1 {
		return []int{1}
	}
	res := make([]int, 0)
	ary := [26]int{}
	for i, v := range s {
		ary[v-'a'] = i
	}
	lastIdx := 0
	minLen := ary[s[0]-'a']
	for curIdx := 0; curIdx < len(s); curIdx++ {
		minLen = Max(ary[s[curIdx]-'a'], minLen)
		if curIdx == minLen {
			res = append(res, curIdx-lastIdx+1)
			lastIdx = curIdx + 1
		}
	}

	return res
}
