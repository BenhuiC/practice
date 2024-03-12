package array

func shiftingLetters(s string, shifts []int) string {
	n := len(s)
	if n < 1 {
		return s
	}
	sum := 0
	for _, v := range shifts {
		sum += v
	}

	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = 'a' + (s[i]-'a'+byte(sum%26))%26
		sum -= shifts[i]
	}

	return string(res)
}

func shiftingLetters2(s string, shifts []int) string {
	n := len(s)
	if n < 1 {
		return s
	}
	sum := 0

	res := make([]byte, n)
	for i := n - 1; i >= 0; i-- {
		sum += shifts[i]
		res[i] = 'a' + (s[i]-'a'+byte(sum%26))%26
	}

	return string(res)
}
