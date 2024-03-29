package array

func sumSubarrayMins(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	leftMinCnt := make([]int, n)
	rightMinCnt := make([]int, n)
	for i := 0; i < n; i++ {
		l, r := i, i
		for l > 0 && arr[l-1] >= arr[i] {
			l--
		}
		for r < n-1 && arr[r+1] >= arr[i] {
			r++
		}
		leftMinCnt[i] = i - l
		rightMinCnt[i] = r - i
	}

	mod := 1000000007
	res := 0
	for i := 0; i < n; i++ {
		mul := leftMinCnt[i] + rightMinCnt[i] + 1 + leftMinCnt[i]*rightMinCnt[i]
		res += mul * arr[i]
		res = res % mod
	}
	return res
}
