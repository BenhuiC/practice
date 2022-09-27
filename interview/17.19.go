package interview

func missingTwo(nums []int) []int {
	n := len(nums) + 2
	x := 0
	for _, v := range nums {
		x = x ^ v
	}
	for i := 1; i <= n; i++ {
		x = x ^ i
	}
	lsp := x & -x
	type1, type2 := 0, 0
	for _, v := range nums {
		if v&lsp == 0 {
			type1 = type1 ^ v
		} else {
			type2 = type2 ^ v
		}
	}
	for i := 1; i <= n; i++ {
		if i&lsp == 0 {
			type1 = type1 ^ i
		} else {
			type2 = type2 ^ i
		}
	}

	res := []int{type1, type2}
	return res
}
