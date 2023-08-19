package partice

func maxChunksToSorted(arr []int) int {
	res := 0
	var left, right, i int
	for {
		if left >= len(arr) {
			break
		}
		right = arr[left]
		for i = left; i <= right && i < len(arr); i++ {
			if arr[i] > right {
				right = arr[i]
			}
		}
		res++
		left = right + 1
	}

	return res
}
