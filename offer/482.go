package offer

func licenseKeyFormatting(s string, k int) string {
	byteAry := make([]byte, 0)
	curLen := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '-' {
			continue
		}
		curLen++
		c := s[i]
		if c >= 'a' && c <= 'z' {
			c = c - 'a' + 'A'
		}
		tmp := []byte{c}
		if curLen == k && i != 0 {
			tmp = append([]byte{'-'}, tmp...)
			curLen = 0
		}
		byteAry = append(tmp, byteAry...)
	}
	if len(byteAry) > 0 && byteAry[0] == '-' {
		byteAry = byteAry[1:]
	}

	return string(byteAry)
}
