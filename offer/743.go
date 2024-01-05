package offer

import "math"

// Dijkstra
func networkDelayTime(times [][]int, n int, k int) int {
	inf := math.MaxInt >> 1
	nets := make([][]int, n)
	for i := 0; i < n; i++ {
		nets[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i != j {
				nets[i][j] = inf
			}
		}
	}
	for _, t := range times {
		nets[t[0]-1][t[1]-1] = t[2]
	}
	visited := make(map[int]bool)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = inf
	}
	dist[k-1] = 0
	for i := 0; i < n; i++ {
		nextIdx := -1
		dis := inf
		for idx, v := range dist {
			if !visited[idx] && v < dis {
				dis = v
				nextIdx = idx
			}
		}
		if nextIdx == -1 {
			break
		}
		visited[nextIdx] = true
		for j, d := range nets[nextIdx] {
			dist[j] = min(dist[j], dist[nextIdx]+d)
		}
	}

	res := 0
	for _, d := range dist {
		if d == inf {
			return -1
		}
		res = max(res, d)
	}
	return res
}
