package hot

func minWindow(s string, t string) string {
	ls, lt := len(s), len(t)
	if ls < lt {
		return ""
	}

	res := ""
	needMap := make(map[byte]int)
	for _, c := range t {
		needMap[byte(c)]++
	}
	needCnt := lt

	l, r := 0, 0
	for l <= r && r < ls {
		if needCnt > 0 {
			v, ok := needMap[s[r]]
			if ok {
				needMap[s[r]]--
				if v > 0 {
					needCnt--
				}
			}
			r++
		}
		for needCnt == 0 {
			if res == "" || res != "" && r-l < len(res) {
				res = s[l:r]
			}
			if v, ok := needMap[s[l]]; !ok {
				l++
				continue
			} else if v < 0 {
				needMap[s[l]]++
				l++
				continue
			}

			needMap[s[l]]++
			l++
			needCnt++
		}
	}

	return res
}
