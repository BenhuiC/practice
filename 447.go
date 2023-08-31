package partice

import "math"

func numberOfBoomerangs(points [][]int) int {
	res447 := 0
	rg := make([][]float64, len(points))
	for i := range rg {
		rg[i] = make([]float64, len(points))
	}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			l := math.Pow(math.Abs(float64(points[i][0]-points[j][0])), 2) + math.Pow(math.Abs(float64(points[i][1]-points[j][1])), 2)
			rg[i][j] = l
			rg[j][i] = l
		}
	}

	for i := 0; i < len(points); i++ {
		tmp := make(map[float64]int)
		for j := 0; j < len(points); j++ {
			tmp[rg[i][j]]++
		}
		for _, v := range tmp {
			if v > 1 {
				res447 = res447 + v*(v-1)
			}
		}
	}
	return res447
}
