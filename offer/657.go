package offer

func judgeCircle(moves string) bool {
	var i, j int
	for _, c := range moves {
		switch c {
		case 'U':
			i++
		case 'D':
			i--
		case 'L':
			j++
		case 'R':
			j--
		}
	}
	return i == 0 && j == 0
}
