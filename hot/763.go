package hot

func partitionLabels(s string) []int {
	n := len(s)
	ary := make([]int, 26)
	for i, c := range s {
		ary[c-'a'] = i
	}

	res := make([]int, 0)
	var start, end int
	for start < n {
		end = ary[s[start]-'a']
		for i := start; i < end; i++ {
			if ary[s[i]-'a'] > end {
				end = ary[s[i]-'a']
			}
		}
		res = append(res, end-start+1)
		start = end + 1
	}

	return res
}
