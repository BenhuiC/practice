package array

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	res := make([][]int, 0)

	left, right, top, bottom := cStart, cStart, rStart, rStart
	for {
		if top >= 0 {
			for i := left; i <= right && i < cols; i++ {
				if i >= 0 {
					res = append(res, []int{top, i})
				}
			}
		}
		right++

		if right < cols {
			for i := top; i <= bottom && i < rows; i++ {
				if i >= 0 {
					res = append(res, []int{i, right})
				}
			}
		}
		bottom++

		if bottom < rows {
			for i := right; i >= left && i >= 0; i-- {
				if i < cols {
					res = append(res, []int{bottom, i})
				}
			}
		}
		left--

		if left >= 0 {
			for i := bottom; i >= top && i >= 0; i-- {
				if i < rows {
					res = append(res, []int{i, left})
				}
			}
		}
		top--

		if len(res) == rows*cols {
			break
		}
	}

	return res
}
