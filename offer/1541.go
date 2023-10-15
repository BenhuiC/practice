package offer

func minInsertions(s string) int {
	res := 0
	leftNu, rightNu := 0, 0
	nl := len(s)
	for i := nl - 1; i >= 0; {
		for i >= 0 && s[i] == ')' {
			rightNu++
			i--
		}
		for i >= 0 && s[i] == '(' {
			leftNu++
			i--
		}
		needRight := leftNu * 2
		if needRight >= rightNu {
			res += needRight - rightNu
			rightNu = 0
		} else {
			rightNu -= needRight
			if rightNu&1 == 1 {
				res++
				rightNu++
			}
		}
		leftNu = 0
	}
	if rightNu > 0 {
		res += rightNu / 2
	}

	return res
}
