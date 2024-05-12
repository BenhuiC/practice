package array

func prefixesDivBy5(nums []int) []bool {
	n := len(nums)
	if n == 0 {
		return nil
	}
	res := make([]bool, n)
	val := 0
	for i, v := range nums {
		val = val<<1 + v
		if val%5 == 0 {
			res[i] = true
		}
		val = val % 10
	}

	return res
}
