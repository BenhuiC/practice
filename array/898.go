package array

func subarrayBitwiseORs(arr []int) int {
	allMap := make(map[int]struct{})
	lastM := make(map[int]struct{})
	for _, v := range arr {
		allMap[v] = struct{}{}
		next := make(map[int]struct{})
		next[v] = struct{}{}
		for k := range lastM {
			val := k | v
			next[val] = struct{}{}
			allMap[val] = struct{}{}
		}
		lastM = next
	}
	return len(allMap)
}
