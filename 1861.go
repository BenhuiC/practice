package partice

func rotateTheBox(box [][]byte) [][]byte {
	m := len(box)
	if m == 0 {
		return nil
	}
	n := len(box[0])
	res := make([][]byte, n)
	for i := 0; i < n; i++ {
		t := make([]byte, m)
		for i := range t {
			t[i] = '.'
		}
		res[i] = t
	}
	for i := m - 1; i >= 0; i-- {
		targetRow := n - 1
		targetColumn := m - 1 - i
		for j := n - 1; j >= 0 && targetRow >= 0; j-- {
			switch box[i][j] {
			case '#':
				res[targetRow][targetColumn] = '#'
				targetRow--
			case '*':
				res[j][targetColumn] = '*'
				targetRow = j - 1
			case '.':
			}
		}
	}

	return res
}
