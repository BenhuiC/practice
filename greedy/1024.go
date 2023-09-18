package greedy

func videoStitching(clips [][]int, time int) (ans int) {
	maxN := make([]int, time)
	for _, c := range clips {
		if c[0] < time && c[1] > maxN[c[0]] {
			maxN[c[0]] = c[1]
		}
	}
	pre, last := 0, 0
	for i, m := range maxN {
		if m > last {
			last = m
		}
		if i == last {
			return -1
		}
		if i == pre {
			pre = last
			ans++
		}
	}
	return
}
