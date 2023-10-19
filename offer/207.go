package offer

func canFinish(numCourses int, prerequisites [][]int) bool {
	inCnt := make(map[int]int, numCourses)
	postMap := make(map[int][]int, numCourses)
	for _, v := range prerequisites {
		inCnt[v[0]]++
		if postMap[v[1]] == nil {
			postMap[v[1]] = make([]int, 0)
		}
		postMap[v[1]] = append(postMap[v[1]], v[0])
	}

	for {
		find := false
		for k, v := range postMap {
			if inCnt[k] <= 0 {
				for _, p := range v {
					inCnt[p]--
				}
				delete(postMap, k)
				delete(inCnt, k)
				find = true
			}
		}
		if find == false {
			break
		}
	}

	return len(postMap) == 0
}

func canFinish2(numCourses int, prerequisites [][]int) bool {
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
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		numCourses--
		for _, v := range postMap[cur] {
			inCnt[v]--
			if inCnt[v] <= 0 {
				q = append(q, v)
			}
		}
	}

	return numCourses == 0
}
