package array

func minDominoRotations(tops []int, bottoms []int) int {
	n := len(tops)
	if n <= 1 {
		return 0
	}

	top, bottom := 0, 0
	top2B, bot2T := 0, 0
	topVal, bottomVal := tops[0], bottoms[0]
	if topVal != bottomVal {
		top2B++
		bot2T++
	}
	for i := 1; i < n; i++ {
		if bottom == -1 && top == -1 {
			return -1
		}

		if top != -1 {
			if tops[i] != topVal && bottoms[i] != topVal {
				top = -1
				top2B = -1
			} else if topVal != tops[i] {
				top++
			} else if topVal != bottoms[i] {
				top2B++
			}
		}

		if bottom != -1 {
			if tops[i] != bottomVal && bottoms[i] != bottomVal {
				bottom = -1
				bot2T = -1
			} else if bottoms[i] != bottomVal {
				bottom++
			} else if bottomVal != tops[i] {
				bot2T++
			}
		}
	}

	if top == -1 {
		return min(bottom, bot2T)
	} else if bottom == -1 {
		return min(top, top2B)
	}

	top = min(top, top2B)
	bottom = min(bottom, bot2T)
	return min(top, bottom)
}
