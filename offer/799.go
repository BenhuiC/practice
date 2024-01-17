package offer

func champagneTower(poured int, query_row int, query_glass int) float64 {
	preLevel := []float64{float64(poured)}
	for i := 1; i <= query_row; i++ {
		next := make([]float64, i+1)
		for j := range next {
			l := 0.0
			if j-1 >= 0 {
				l = preLevel[j-1] - 1
			}
			if l > 0 {
				next[j] += l / 2
			}

			r := 0.0
			if j < i {
				r = preLevel[j] - 1
			}
			if r > 0 {
				next[j] += r / 2
			}
		}
		preLevel = next
	}
	if preLevel[query_glass] >= 1 {
		return 1
	} else {
		return preLevel[query_glass]
	}
}
