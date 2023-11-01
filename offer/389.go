package offer

func findTheDifference(s string, t string) byte {
	ary := make([]int, 26)
	for i, v := range t {
		ary[v-'a']--
		if i < len(s) {
			ary[s[i]-'a']++
		}
	}
	for i, v := range ary {
		if v == -1 {
			return byte('a' + i)
		}
	}
	return 'a'
}
