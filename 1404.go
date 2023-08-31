package partice

func numSteps(s string) int {
	res := 0
	var carry uint8
	for n := len(s); n > 1; n-- {
		t := s[n-1] - '0' + carry
		switch t {
		case 0:
			res++
		case 1:
			carry = 1
			res += 2
		case 2:
			carry = 1
			res += 1
		}
	}
	if carry == 1 {
		res += 1
	}

	return res
}
