package offer

func reorganizeString(s string) string {
	ch := make([]int, 26)
	n := len(s)
	maxCnt := 0
	// 统计每个字母出现次数
	for _, c := range s {
		idx := c - 'a'
		ch[idx]++
		if ch[idx] > maxCnt {
			maxCnt = ch[idx]
		}
	}
	// 如果字母的最大出现次数大于(n+1)/2则比有相邻的字母
	if maxCnt > (n+1)/2 {
		return ""
	}
	res := make([]byte, n)
	oddIdx, evenIdx, half := 0, 1, n/2
	for i, c := range ch {
		cha := 'a' + i
		// 先尝试放偶数位，偶数为放完后再放奇数位（如果偶数位放完仍有剩余，则说明剩下的字母出现次数必小于等于n/2）
		for c > 0 && c <= half && evenIdx < n {
			res[evenIdx] = byte(cha)
			c--
			evenIdx += 2
		}
		// 字母次数大于half的放在奇数位
		for c > 0 && oddIdx < n {
			res[oddIdx] = byte(cha)
			c--
			oddIdx += 2
		}
	}

	return string(res)
}
