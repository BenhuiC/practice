package partice

import "sort"

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	trimLen := len(arr) / 20
	arr = arr[trimLen : len(arr)-trimLen]
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return float64(sum) / float64(len(arr))
}
