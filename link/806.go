package link

func numberOfLines(widths []int, s string) []int {
	if len(s) == 0 {
		return []int{0, 0}
	}
	lineCnt := 1
	curCnt := 0
	for _, c := range s {
		curW := widths[c-'a']
		if curCnt+curW <= 100 {
			curCnt += curW
		} else {
			lineCnt++
			curCnt = curW
		}
	}
	return []int{lineCnt, curCnt}
}
