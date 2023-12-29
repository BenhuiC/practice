package offer

func countBinarySubstrings(s string) int {
	n := len(s)
	var res, lastOneCnt int
	var i int
	for i < n {
		zeroCnt := 0
		var j int
		for j = i; j < n && s[j] == '0'; j++ {
			zeroCnt++
		}
		res += min(zeroCnt, lastOneCnt)
		oneCnt := 0
		for j < n && s[j] == '1' {
			j++
			oneCnt++
		}
		res += min(zeroCnt, oneCnt)

		lastOneCnt = oneCnt
		i = j
	}

	return res
}
