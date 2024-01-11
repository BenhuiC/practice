package offer

func canTransform(start string, end string) bool {
	if start == end {
		return true
	}
	ls, le := len(start), len(end)
	if ls != le {
		return false
	}

	i1, i2 := 0, 0
	for i1 < ls && i2 < le {
		for i1 < ls && start[i1] == 'X' {
			i1++
		}
		for i2 < le && end[i2] == 'X' {
			i2++
		}
		if i1 < ls && i2 < le {
			if start[i1] != end[i2] {
				return false
			}
			if start[i1] == 'L' && i1 < i2 || start[i1] == 'R' && i1 > i2 {
				return false
			}
			i1++
			i2++
		}
	}
	for i1 < ls {
		if start[i1] != 'X' {
			return false
		}
		i1++
	}
	for i2 < le {
		if end[i2] != 'X' {
			return false
		}
		i2++
	}

	return true
}
