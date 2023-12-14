package offer

import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	l1Map := make(map[string]int)
	for i, v := range list1 {
		l1Map[v] = i
	}
	res := make([]string, 0)
	maxSum := math.MaxInt
	for i, v := range list2 {
		if i > maxSum {
			break
		}
		if j, ok := l1Map[v]; ok {
			if i+j < maxSum {
				maxSum = i + j
				res = []string{v}
			} else if i+j == maxSum {
				res = append(res, v)
			}
		}
	}
	return res
}
