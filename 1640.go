package partice

func canFormArray(arr []int, pieces [][]int) bool {
	res := true
	m := make(map[int]int)
	for i, v := range pieces {
		if len(v) == 0 {
			continue
		}
		m[v[0]] = i
	}
	for i := 0; i < len(arr); {
		idx, ok := m[arr[i]]
		if !ok {
			res = false
			break
		}
		for _, v := range pieces[idx] {
			if v != arr[i] {
				return false
			}
			i++
		}
	}

	return res
}
