package partice

func canBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
	res := true
	m := make(map[int]int)
	for i := 0; i < len(target); i++ {
		m[target[i]]++
		m[arr[i]]--
	}
	for _, v := range m {
		if v != 0 {
			res = false
			break
		}
	}

	return res
}
