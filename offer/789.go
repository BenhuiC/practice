package offer

func escapeGhosts(ghosts [][]int, target []int) bool {
	dis := disOf2(0, 0, target[0], target[1])
	for _, g := range ghosts {
		if disOf2(g[0], g[1], target[0], target[1]) <= dis {
			return false
		}
	}
	return true
}

func disOf2(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}
