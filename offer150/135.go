package offer150

func candy(ratings []int) int {
	n := len(ratings)
	arr := make([]int, len(ratings))
	res := 0
	for i := range ratings {
		if i == 0 || ratings[i] <= ratings[i-1] {
			arr[i] = 1
		} else {
			arr[i] = arr[i-1] + 1
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i == n-1 || ratings[i] <= ratings[i+1] {
			arr[i] = max(arr[i], 1)
		} else {
			arr[i] = max(arr[i], arr[i+1]+1)
		}
		res += arr[i]
	}
	return res
}
