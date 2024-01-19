package offer

func soupServings(n int) float64 {
	n = (n + 24) / 25
	if n >= 179 {
		return 1
	}

	type item struct {
		a, b int
		x    float64
	}

	var proba, probb float64
	dec := [][2]int{{4, 0}, {3, 1}, {2, 2}, {1, 3}}

	ary := make([]item, 0)
	ary = append(ary, item{n, n, 1})
	for len(ary) > 0 {
		cur := ary[0]
		ary = ary[1:]

		for _, d := range dec {
			a, b := cur.a-d[0], cur.b-d[1]
			if a > 0 && b > 0 {
				x := cur.x / 4
				// TODO oom
				ary = append(ary, item{a, b, x})
			} else if a <= 0 && b > 0 {
				proba += cur.x / 4
			} else if a <= 0 && b <= 0 {
				probb += cur.x / 4
			}

			//fmt.Println(len(ary))
		}
	}

	return proba + probb/2
}
