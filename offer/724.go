package offer

func pivotIndex(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	s := 0
	for i, n := range nums {
		if s<<1+n == sum {
			return i
		}
		s += n
	}
	return -1
}
