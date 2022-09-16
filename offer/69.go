package offer

func peakIndexInMountainArray(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}
	max := arr[0]
	resOffer69 := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			resOffer69 = i
		}
	}
	return resOffer69
}
