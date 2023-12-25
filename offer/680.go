package offer

func validPalindrome(s string) bool {
	isPalindrome := func(str string) bool {
		l, r := 0, len(str)-1
		for l < r && str[l] == str[r] {
			l++
			r++
		}
		return l >= r
	}
	l, r := 0, len(s)-1
	for l < r {
		if s[l] == s[r] {
			l++
			r--
			continue
		}
		return isPalindrome(s[l:r]) || isPalindrome(s[l+1:r+1])
	}
	return true
}
