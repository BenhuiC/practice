package offer

func isAnagram(s string, t string) bool {
	sl, tl := len(s), len(t)
	if sl != tl {
		return false
	}
	sary := [26]int{}
	tary := [26]int{}
	for i := 0; i < sl; i++ {
		sary[s[i]-'a']++
		tary[t[i]-'a']++
	}
	return sary == tary
}
