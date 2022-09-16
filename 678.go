package partice

func checkValidString(s string) bool {
	xary, lary := []int{}, []int{}
	for i, v := range s {
		if v == '(' {
			lary = append(lary, i)
			continue
		} else if v == '*' {
			xary = append(xary, i)
			continue
		}
		if len(lary) > 0 {
			lary = lary[:len(lary)-1]
			continue
		}
		if len(xary) > 0 {
			xary = xary[:len(xary)-1]
			continue
		}
		return false
	}
	if len(xary) < len(lary) {
		return false
	}
	for len(xary) > 0 && len(lary) > 0 {
		if lary[len(lary)-1] < xary[len(xary)-1] {
			lary = lary[:len(lary)-1]
			xary = xary[:len(xary)-1]
			continue
		}
		break
	}
	if len(lary) > 0 {
		return false
	}
	return true
}
