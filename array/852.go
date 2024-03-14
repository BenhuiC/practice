package array

func peakIndexInMountainArray(arr []int) int {
	n := len(arr)
	if n < 1 {
		return 0
	}
	l, r := 0, n-1
	for l < r {
		mid := l + (r-l)/2
		if mid == 0 {
			l++
			continue
		} else if mid == n-1 {
			r--
			continue
		}
		if arr[mid+1] < arr[mid] && arr[mid-1] < arr[mid] {
			return mid
		} else if arr[mid+1] >= arr[mid] {
			l++
		} else if arr[mid-1] >= arr[mid] {
			r--
		}
	}
	return 0
}
