package partice

func peopleIndexes(favoriteCompanies [][]string) []int {
	res1452 := make([]int, 0)
	buArr := make([]bool, len(favoriteCompanies))
	for i := 0; i < len(favoriteCompanies)-1; i++ {
		for j := i + 1; j < len(favoriteCompanies); j++ {
			if contain(favoriteCompanies[i], favoriteCompanies[j]) {
				buArr[j] = true
			}
			if contain(favoriteCompanies[j], favoriteCompanies[i]) {
				buArr[i] = true
			}
		}
	}
	for i, v := range buArr {
		if !v {
			res1452 = append(res1452, i)
		}
	}

	return res1452
}

func contain(a, b []string) bool {
	mapA := make(map[string]struct{})
	for _, v := range a {
		mapA[v] = struct{}{}
	}
	for _, v := range b {
		if _, ok := mapA[v]; !ok {
			return false
		}
	}
	return true
}
