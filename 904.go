package partice

func totalFruit(fruits []int) int {
	res := 0
	if len(fruits) < 3 {
		return len(fruits)
	}
	kindA, kindB := -1, -1
	idxA, idxB := 0, 0
	lastA, lastB := 0, 0
	left := 0

	for right, v := range fruits {
		if kindA == -1 || kindA == v {
			if kindA == -1 {
				kindA = v
				idxA = right
			}
			lastA = right
		} else if kindB == -1 || kindB == v {
			if kindB == -1 {
				kindB = v
				idxB = right
			}
			lastB = right
		} else {
			if lastA > lastB { //保留A
				left = Max(idxA, lastB+1)
				kindB = v
				idxB = right
				lastB = right
			} else {
				left = Max(idxB, lastA+1)
				kindA = v
				idxA = right
				lastA = right
			}
		}

		res = Max(res, right-left+1)
	}

	return res
}
