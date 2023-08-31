package partice

func findDuplicate(nums []int) int {
	res287 := 0
	ary := make([]int, len(nums))
	for _, v := range nums {
		if ary[v] != 0 {
			res287 = v
			break
		}
		ary[v] = 1
	}

	return res287
}
