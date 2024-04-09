package array

func sortArray(nums []int) []int {
	n := len(nums)

	return mergeSort(nums, 0, n)
}

func quickSort(arr []int, l, r int) []int {
	if l >= r || r-l < 2 {
		return arr
	}
	pos := arr[l]
	idx := l
	for i := l; i < r; i++ {
		v := arr[i]
		if v <= pos {
			arr[idx], arr[i] = arr[i], arr[idx]
			idx++
		}
	}
	arr[idx-1], arr[l] = arr[l], arr[idx-1]

	quickSort(arr, l, idx-1)
	quickSort(arr, idx, r)
	return arr
}

func mergeSort(arr []int, l, r int) []int {
	if l >= r || r-l < 2 {
		return arr[l:r]
	}
	mid := l + (r-l)/2
	left := mergeSort(arr, l, mid)
	right := mergeSort(arr, mid, r)

	res := make([]int, 0, r-l)
	var idx1, idx2 int

	for idx1 < len(left) && idx2 < len(right) {
		if left[idx1] < right[idx2] {
			res = append(res, left[idx1])
			idx1++
		} else {
			res = append(res, right[idx2])
			idx2++
		}
	}
	for idx1 < len(left) {
		res = append(res, left[idx1])
		idx1++
	}
	for idx2 < len(right) {
		res = append(res, right[idx2])
		idx2++
	}

	return res
}
