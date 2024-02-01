package hot

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for cur, v := range nums {
		if last, ok := m[target-v]; ok {
			return []int{last, cur}
		}
		m[v] = cur
	}
	return nil
}
