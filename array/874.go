package array

func robotSim(commands []int, obstacles [][]int) int {
	type pair struct {
		x, y int
	}
	obsMap := make(map[pair]bool)
	for _, o := range obstacles {
		obsMap[pair{x: o[0], y: o[1]}] = true
	}
	dirStep := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	dir := 0
	var x, y int
	var res int
	for _, c := range commands {
		if c == -2 {
			dir = (dir + 3) % 4
			continue
		}
		if c == -1 {
			dir = (dir + 1) % 4
			continue
		}
		step := dirStep[dir]
		for i := 0; i < c; i++ {
			tarX, tarY := x+step[0], y+step[1]
			if obsMap[pair{x: tarX, y: tarY}] {
				break
			}
			x, y = tarX, tarY
		}
		res = max(res, x*x+y*y)
	}

	return res
}
