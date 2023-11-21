package offer

import "math/rand"

func rand10() int {
	var first, second int
	for {
		first = rand7()
		if first <= 6 {
			break
		}
	}
	for {
		second = rand7()
		if second <= 5 {
			break
		}
	}
	if first&1 == 1 {
		return second
	}
	return 5 + second
}

func rand7() int {
	return rand.Intn(7) + 1
}
