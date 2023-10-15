package offer

func numMatchingSubseq(s string, words []string) int {
	res := 0
	wordMap := make(map[string]int)
	for _, v := range words {
		wordMap[v]++
	}
	for k, v := range wordMap {
		if isSubsequence(k, s) {
			res += v
		}
	}

	return res
}
