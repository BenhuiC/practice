package array

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var f func(i int) (re, left, right int)
	f = func(i int) (re, left, right int) {
		left = i
		for right = i; right < len(s)-1 && s[right+1] == s[right]; right++ {
		}
		re = right + 1
		for left > 0 && right < len(s)-1 && s[left-1] == s[right+1] {
			left--
			right++
		}

		return
	}
	res := s[:1]
	for i := 0; i < len(s); {
		t, l, r := f(i)
		if r-l+1 > len(res) {
			res = s[l : r+1]
		}
		i = t
	}

	return res
}
