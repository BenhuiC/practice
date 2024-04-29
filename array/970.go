package array

func powerfulIntegers(x int, y int, bound int) []int {
	m := make(map[int]bool)
	res := make([]int, 0)
	i, j := 1, 1
	for i < bound {
		for j = 1; j < bound; j *= y {
			tmp := i + j
			if m[tmp] {
				continue
			}
			if tmp <= bound {
				res = append(res, tmp)
				m[tmp] = true
			} else {
				break
			}
			if y == 1 {
				break
			}
		}

		i *= x
		if i == 1 {
			break
		}
	}

	return res
}
