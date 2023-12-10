package offer

func checkRecord(s string) bool {
	aCnt, lCnt := 0, 0
	for _, v := range s {
		if v == 'P' {
			lCnt = 0
		} else if v == 'A' {
			aCnt++
			lCnt = 0
		} else {
			lCnt++
		}
		if aCnt >= 2 || lCnt >= 3 {
			return false
		}
	}
	return true
}
