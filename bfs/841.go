package bfs

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	if n < 1 {
		return true
	}
	m := make(map[int]bool)
	curKeys := rooms[0]
	m[0] = true
	for len(curKeys) > 0 {
		nextKeys := make([]int, 0)
		for _, v := range curKeys {
			if !m[v] {
				nextKeys = append(nextKeys, rooms[v]...)
				m[v] = true
			}
		}
		curKeys = nextKeys
	}
	return len(m) == n
}
