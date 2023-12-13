package offer

import (
	"math"
	"sort"
)

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	l := []float64{dis(p1, p2), dis(p1, p3), dis(p1, p4), dis(p2, p3), dis(p2, p4), dis(p3, p4)}
	sort.Float64s(l)
	return l[0] != 0 && l[0] == l[3] && l[4] == l[5] && l[0]*2 == l[5]
}

func dis(p1, p2 []int) float64 {
	return math.Pow(float64(p1[0]-p2[0]), 2) + math.Pow(float64(p1[1]-p2[1]), 2)
}
