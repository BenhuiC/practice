package array

func findJudge(n int, trust [][]int) int {
	in, out := make(map[int]int), make(map[int]int)
	for _, t := range trust {
		o, i := t[0], t[1]
		in[i]++
		out[o]++
	}
	res := -1
	for i := 1; i <= n; i++ {
		if in[i] == n-1 && out[i] == 0 {
			if res != -1 {
				return -1
			}
			res = i
		}
	}

	return res
}
