package partice

func merge56(intervals [][]int) [][]int {
	result := make([][]int, 0)
	if len(intervals) <= 1 {
		return intervals
	}
	for i := 0; i < len(intervals); i++ {
		if len(result) == 0 {
			result = append(result, intervals[i])
			continue
		}
		for j := 0; j < len(result); j++ {
			if v, ok := merge2ary(result[j], intervals[i]); ok {

			}
		}
	}

	return result
}

func merge2ary(a, b []int) ([]int, bool) {
	return nil, false
}
