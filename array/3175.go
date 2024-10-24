package array

func findWinningPlayer(skills []int, k int) int {
	n := len(skills)
	q := make([]int, n)
	maxVal := 0
	maxIdx := 0
	for i := 0; i < n; i++ {
		q[i] = i
		if skills[i] > maxVal {
			maxIdx = i
			maxVal = skills[i]
		}
	}
	if k >= n {
		return maxIdx
	}
	idx := 0
	win := 0
	for win < k {
		idxA := idx
		idxB := (idx + 1) % n
		valA := skills[q[idxA]]
		valB := skills[q[idxB]]
		if valA > valB {
			q[idxA], q[idxB] = q[idxB], q[idxA]
			win++
		} else {
			win = 1
		}
		idx = idxB
	}

	return q[idx]
}
