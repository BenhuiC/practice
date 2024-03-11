package array

func longestMountain(arr []int) int {
	n := len(arr)
	res := 0
	l, r := 0, 0
	for l < n {
		r = l
		inc, dec := false, false
		for r < n-1 && arr[r+1] > arr[r] {
			r++
			inc = true
		}
		for r < n-1 && arr[r+1] < arr[r] {
			r++
			dec = true
		}
		if inc && dec && r-l+1 >= 3 {
			res = max(res, r-l+1)
		}
		l = r

		if l == n-1 || arr[l+1] == arr[l] {
			l++
		}
	}
	return res
}
