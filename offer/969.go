package offer

func pancakeSort(arr []int) []int {
	res := make([]int, 0)
	nl := len(arr)
	for n := nl; n > 0; n-- {
		idx := 0
		for j := 1; j < n; j++ {
			if arr[j] > arr[idx] {
				idx = j
			}
		}
		if idx == n-1 {
			continue
		}

		for i := 0; i < (idx+1)/2; i++ {
			arr[i], arr[idx-i] = arr[idx-i], arr[i]
		}
		for i := 0; i < n/2; i++ {
			arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
		}

		res = append(res, idx+1, n)
	}

	return res
}
