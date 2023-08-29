package array

func minWindow(s string, t string) string {
	slen, tlen := len(s), len(t)
	if tlen > slen {
		return ""
	}
	res := ""
	need := tlen
	needMap := make(map[byte]int)
	for i := range t {
		needMap[t[i]] += 1
	}
	// l,r为窗口的左右两侧
	l, r := 0, 0
	for l <= r && r < slen {
		if need > 0 {
			if v, ok := needMap[s[r]]; ok {
				needMap[s[r]]--
				if v > 0 {
					need--
				}
			}
			r++
		}
		for need == 0 {
			// need=0,说明l到r包含所有t中字符
			if res == "" || res != "" && r-l < len(res) {
				res = s[l:r]
			}
			// 左侧有多余字符，窗口向右缩
			if v, ok := needMap[s[l]]; !ok || v < 0 {
				if v < 0 {
					needMap[s[l]]++
				}
				l++
				continue
			}
			// 此时窗口中包含所有t中的字符，并且左则不含有多余的字符，将左侧窗口右移，使窗口可以继续向右拓展
			needMap[s[l]]++
			need++
			l++
		}
	}

	return res
}
