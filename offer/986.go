package offer

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var idx1, idx2 int
	res := make([][]int, 0)
	l1, l2 := len(firstList), len(secondList)
	for idx1 < l1 && idx2 < l2 {
		if firstList[idx1][0] > secondList[idx2][1] {
			idx2++
			continue
		} else if secondList[idx2][0] > firstList[idx1][1] {
			idx1++
			continue
		}
		var l, r int
		if firstList[idx1][0] >= secondList[idx2][0] {
			l = firstList[idx1][0]
			if secondList[idx2][1] > firstList[idx1][1] {
				r = firstList[idx1][1]
				idx1++
			} else {
				r = secondList[idx2][1]
				idx2++
			}
		} else if secondList[idx2][0] >= firstList[idx1][0] {
			l = secondList[idx2][0]
			r = firstList[idx1][1]
			if firstList[idx1][1] > secondList[idx2][1] {
				r = secondList[idx2][1]
				idx2++
			} else {
				r = firstList[idx1][1]
				idx1++
			}
		}

		res = append(res, []int{l, r})
	}
	return res
}
