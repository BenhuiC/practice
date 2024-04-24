package array

import (
	"sort"
)

func canReorderDoubled(arr []int) bool {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	sort.Slice(arr, func(i, j int) bool {
		a, b := arr[i], arr[j]
		if a < 0 {
			a = -a
		}
		if b < 0 {
			b = -b
		}
		return a < b
	})
	for _, v := range arr {
		if m[v] <= 0 {
			continue
		}
		n := m[v*2]
		if n <= 0 {
			return false
		}
		m[v]--
		m[v*2]--
	}

	return true
}
