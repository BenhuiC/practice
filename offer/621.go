package offer

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	cntMap := make(map[byte]int)
	for _, v := range tasks {
		cntMap[v]++
	}
	maxExec, maxExecCnt := 0, 0
	for _, v := range cntMap {
		if v > maxExec {
			maxExec = v
			maxExecCnt = 1
		} else if maxExec == v {
			maxExecCnt++
		}
	}
	res := (maxExec-1)*(n+1) + maxExecCnt
	if res > len(tasks) {
		return res
	}

	return len(tasks)
}
