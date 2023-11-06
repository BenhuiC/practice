package offer

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] != people[j][0] {
			return people[i][0] < people[j][0]
		}
		return people[i][1] > people[j][1]
	})
	n := len(people)
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		item := people[i]
		space := item[1] + 1
		for j := 0; j < n; j++ {
			if res[j] == nil {
				space--
				if space == 0 {
					res[j] = item
					break
				}
			}
		}
	}

	return res
}
