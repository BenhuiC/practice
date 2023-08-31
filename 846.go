package partice

import "math"

func isNStraightHand(hand []int, groupSize int) bool {
	m := make(map[int]int)
	min := math.MaxInt32
	for _, v := range hand {
		m[v]++
		if v < min {
			min = v
		}
	}
	for len(m) != 0 {
		if min == -1 {
			min = math.MaxInt32
			for k := range m {
				if k < min {
					min = k
				}
			}
		}
		for i := min; i < min+groupSize; i++ {
			v, ok := m[i]
			if !ok || v == 0 {
				return false
			}
			if v == 1 {
				delete(m, i)
			} else {
				m[i]--
			}
		}
		min = -1
	}
	return true
}
