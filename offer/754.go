package offer

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	k := 0
	for target > 0 {
		k++
		target -= k
	}
	if target%2 == 0 {
		return k
	}
	return k + 1 + k%2
}

func reachNumber2(target int) int {
	step := 1
	ary := []int{0}
	for len(ary) != 0 {
		next := []int{}
		for _, v := range ary {
			i1, i2 := v-step, v+step
			if i1 == target || i2 == target {
				return step
			}
			next = append(next, i1, i2)
		}
		step++
		ary = next
	}

	return -1
}
