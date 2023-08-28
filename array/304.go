package array

type NumMatrix struct {
	prefixSum [][]int
}

func Constructor2(matrix [][]int) NumMatrix {
	ary := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		t := make([]int, len(matrix[0]))
		s := 0
		for j := 0; j < len(matrix[i]); j++ {
			s += matrix[i][j]
			t[j] = s
			if i != 0 {
				t[j] += ary[i-1][j]
			}
		}
		ary[i] = t
	}
	return NumMatrix{prefixSum: ary}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {

	a, b, c := 0, 0, 0
	if row1 != 0 {
		b = this.prefixSum[row1-1][col2]
	}
	if col1 != 0 {
		c = this.prefixSum[row2][col1-1]
	}
	if row1 != 0 && col1 != 0 {
		a = this.prefixSum[row1-1][col1-1]
	}

	res := this.prefixSum[row2][col2] + a - b - c
	return res
}
