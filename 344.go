package partice

func reverseString(s []byte) {
	if len(s) <= 1 {
		return
	}
	revers(s, 0, len(s)-1)
	return
}

func revers(s []byte, l, h int) {
	if l >= h {
		return
	}
	s[l], s[h] = s[h], s[l]
	revers(s, l+1, h-1)
}
