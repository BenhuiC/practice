package offer

import "fmt"

func movingCount(m int, n int, k int) int {
	step := [][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	mp := make(map[string]bool)
	res13 := 1
	que := make([][2]int, 0)
	que = append(que, [2]int{0, 0})
	mp["0,0"] = true
	for len(que) != 0 {
		st := que[0]
		que = que[1:]
		for _, v := range step {
			tmpx := st[0] + v[0]
			tmpy := st[1] + v[1]
			if tmpx < 0 || tmpy < 0 {
				continue
			}
			if tmpx >= m || tmpy >= n {
				continue
			}
			ky := fmt.Sprintf("%d,%d", tmpx, tmpy)
			if mp[ky] {
				continue
			}
			if sm(tmpx, tmpy) <= k {
				que = append(que, [2]int{tmpx, tmpy})
				mp[ky] = true
				res13++
			}
		}
	}

	return res13
}

func sm(i, j int) int {
	s := 0
	for i != 0 {
		s += i % 10
		i /= 10
	}
	for j != 0 {
		s += j % 10
		j /= 10
	}
	return s
}
