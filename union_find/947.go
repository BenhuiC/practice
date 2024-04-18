package union_find

func removeStones(stones [][]int) int {
	count := 0
	parents := make(map[int]int)
	var find func(x int) int
	find = func(x int) int {
		if v, ok := parents[x]; !ok {
			parents[x] = x
			count++
			return x
		} else if v != x {
			parents[x] = find(parents[x])
		}
		return parents[x]
	}
	union := func(x, y int) {
		parentX := find(x)
		parentY := find(y)
		if parentX == parentY {
			return
		}
		// 当parentX！=parentY时，count++过两次，这里需要减1
		parents[parentX] = parentY
		count--
	}
	for _, stone := range stones {
		union(stone[0]+10001, stone[1])
	}
	return len(stones) - count
}
