package offer

func findOrder(numCourses int, prerequisites [][]int) []int {
	inCnt := make(map[int]int, numCourses)
	postMap := make(map[int][]int, numCourses)
	for _, v := range prerequisites {
		inCnt[v[0]]++
		if postMap[v[1]] == nil {
			postMap[v[1]] = make([]int, 0)
		}
		postMap[v[1]] = append(postMap[v[1]], v[0])
	}

	q := make([]int, 0, numCourses)
	for i := 0; i < numCourses; i++ {
		if inCnt[i] == 0 {
			q = append(q, i)
		}
	}
	res := make([]int, 0, numCourses)
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		numCourses--
		res = append(res, cur)
		for _, v := range postMap[cur] {
			inCnt[v]--
			if inCnt[v] <= 0 {
				q = append(q, v)
			}
		}
	}
	if numCourses == 0 {
		return res
	}
	return nil
}
