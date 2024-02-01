package hot

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, str := range strs {
		ary := [26]int{}
		for _, c := range str {
			ary[c-'a']++
		}
		m[ary] = append(m[ary], str)
	}
	res := make([][]string, 0, len(m))
	for k := range m {
		res = append(res, m[k])
	}
	return res
}
