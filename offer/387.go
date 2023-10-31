package offer

func firstUniqChar(s string) int {
	mp := make(map[int32]int)
	for i, v := range s {
		if mp[v] == -1 {
			continue
		}
		if _, ok := mp[v]; !ok {
			mp[v] = i
		} else {
			mp[v] = -1
		}
	}
	res := -1
	for _, v := range mp {
		if v >= 0 && (res == -1 || v < res) {
			res = v
		}
	}
	return res
}
