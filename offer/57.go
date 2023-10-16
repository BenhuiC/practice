package offer

func insert(intervals [][]int, newInterval []int) [][]int {
	nl := len(intervals)
	res := make([][]int, 0, nl+1)
	if nl == 0 {
		res = append(res, newInterval)
		return res
	}
	if newInterval[0] > intervals[nl-1][1] {
		intervals = append(intervals, newInterval)
		return intervals
	}
	for idx, v := range intervals {
		if newInterval[0] > v[1] {
			res = append(res, intervals[idx])
			continue
		}
		if newInterval[1] < v[0] {
			res = append(res, newInterval)
			res = append(res, intervals[idx:]...)
			return res
		}
		cur := make([]int, 2)
		cur[0], cur[1] = v[0], v[1]
		if newInterval[0] < v[0] {
			cur[0] = newInterval[0]
		}
		if newInterval[1] > v[1] {
			cur[1] = newInterval[1]
		}
		j := idx + 1
		for j < nl && intervals[j][0] <= cur[1] {
			if intervals[j][1] > cur[1] {
				cur[1] = intervals[j][1]
			}
			j++
		}
		res = append(res, cur)
		res = append(res, intervals[j:]...)
		break
	}

	return res
}
