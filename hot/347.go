package hot

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	buckets := make([][]int, len(nums)+1)
	for key, val := range m {
		if buckets[val] == nil {
			buckets[val] = make([]int, 0)
		}
		buckets[val] = append(buckets[val], key)
	}

	res := make([]int, 0, k)
	for i := len(nums); i >= 0; i-- {
		if k >= len(buckets[i]) {
			res = append(res, buckets[i]...)
			k -= len(buckets[i])
		} else {
			res = append(res, buckets[i][:k]...)
			k = 0
		}

		if k == 0 {
			break
		}
	}

	return res
}
