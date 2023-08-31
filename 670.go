package partice

import "fmt"

func maximumSwap(num int) int {
	if num < 10 {
		return num
	}
	s := fmt.Sprint(num)
	index := 0
	max := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] >= max {
			max = s[i]
			index = i
		}
	}
	if index == 0 {
		return num
	}
	add := int(max - s[0])
	sub := int(max - s[0])
	for i := 1; i < len(s); i++ {
		add *= 10
		if i < len(s)-index {
			sub *= 10
		}
	}
	num += add
	num -= sub

	return num
}
