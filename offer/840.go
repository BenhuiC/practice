package offer

func numMagicSquaresInside(grid [][]int) int {
	m := len(grid)
	if m < 3 {
		return 0
	}
	n := len(grid)
	if n < 3 {
		return 0
	}

	sq := func(n ...int) bool {
		if len(n) != 9 {
			return false
		}
		mi, mx := 9, 0
		for _, v := range n {
			if v < mi {
				mi = v
			}
			if v > mx {
				mx = v
			}
		}
		if mi != 1 || mx != 9 {
			return false
		}
		return n[0]+n[1]+n[2] == 15 &&
			n[3]+n[4]+n[5] == 15 &&
			n[6]+n[7]+n[8] == 15 &&
			n[0]+n[3]+n[6] == 15 &&
			n[1]+n[4]+n[7] == 15 &&
			n[2]+n[5]+n[8] == 15 &&
			n[0]+n[4]+n[8] == 15 &&
			n[6]+n[4]+n[2] == 15
	}

	var res int

	for i := 0; i < m-2; i++ {
		for j := 0; j < n-2; j++ {
			if grid[i+1][j+1] != 5 {
				continue
			}
			if sq(grid[i][j], grid[i][j+1], grid[i][j+2], grid[i+1][j], grid[i+1][j+1], grid[i+1][j+2], grid[i+2][j], grid[i+2][j+1], grid[i+2][j+2]) {
				res++
			}
		}
	}

	return res
}
