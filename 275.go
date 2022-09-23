package partice

func hIndex(citations []int) int {
	n := len(citations)
	res := 0
	left, right := 0, len(citations)
	for left < right {
		mid := (left + right) >> 1
		sub := n - mid
		if citations[mid] >= sub {
			res = Max(res, sub)
			right = mid
			if citations[mid] == sub {
				break
			}
		} else {
			left = mid + 1
		}
	}
	return res
}
