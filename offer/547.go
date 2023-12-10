package offer

func findCircleNum(isConnected [][]int) int {
	res := 0
	n := len(isConnected)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if isConnected[i][j] == 0 {
				continue
			}
			res++
			q := []int{j}
			for len(q) != 0 {
				t := q[0]
				q = q[1:]
				isConnected[t][t] = 0
				for idx, nex := range isConnected[t] {
					if nex == 0 {
						continue
					}
					isConnected[t][idx] = 0
					isConnected[idx][t] = 0
					q = append(q, idx)
				}
			}
		}
	}

	return res
}
