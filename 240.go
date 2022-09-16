package partice

func searchMatrix240(matrix [][]int, target int) bool {
	if target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		l, r := 0, n-1
		for l <= r {
			mid := (l + r) / 2
			if matrix[i][mid] > target {
				r = mid - 1
				continue
			}
			if matrix[i][mid] < target {
				l = mid + 1
				continue
			}
			return true
		}
	}
	return false
}
