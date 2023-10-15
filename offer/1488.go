package offer

func avoidFlood(rains []int) []int {
	zeroList := make([]int, 0)
	numCount := make(map[int]int)

	res := make([]int, len(rains))
	for i, v := range rains {
		if v == 0 {
			zeroList = append(zeroList, i)
			res[i] = 1
			continue
		}
		res[i] = -1
		if lastIdx, ok := numCount[v]; !ok {
			numCount[v] = i
		} else {
			zn := len(zeroList)
			if zn == 0 {
				return nil
			}
			// 可以替换为二分查找
			idx := 0
			for idx < zn && zeroList[idx] < lastIdx {
				idx++
			}
			if idx == zn {
				return nil
			}
			res[zeroList[idx]] = v
			zeroList[idx] = -1
			numCount[v] = i
		}
	}

	return res
}
