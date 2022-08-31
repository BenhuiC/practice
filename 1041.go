package partice

func isRobotBounded(instructions string) bool {
	res := false
	p := [2]int{0, 0}
	ary := [4][2]int{
		{0, 1},  // 北
		{1, 0},  // 东
		{0, -1}, // 南
		{-1, 0}, // 西
	}
	l := 0
	for i := range instructions {
		switch instructions[i] {
		case 'G':
			p[0] += ary[l][0]
			p[1] += ary[l][1]
		case 'L':
			l = (l + 3) % 4
		case 'R':
			l = (l + 1) % 4
		}
	}
	if p[0] == 0 && p[1] == 0 {
		res = true
	}
	if l != 0 {
		res = true
	}

	return res
}
