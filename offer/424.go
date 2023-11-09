package offer

func characterReplacement(s string, k int) int {
	n := len(s)
	if n < 2 {
		return n
	}
	res := 0
	l, r := 0, 0
	freq := make([]int, 26)
	mostChartNum := 0
	for r < n {
		freq[s[r]-'A']++
		if freq[s[r]-'A'] > mostChartNum {
			mostChartNum = freq[s[r]-'A']
		}
		r++
		if r-l > mostChartNum+k {
			freq[s[l]-'A']--
			l++
		}
		if r-l > res {
			res = r - l
		}
	}

	return res
}
