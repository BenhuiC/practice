package array

func numSpecialEquivGroups(words []string) int {
	m := make(map[[2][26]int]bool)
	for _, w := range words {
		odd := [26]int{}
		even := [26]int{}
		for i, c := range w {
			if i&1 == 0 {
				odd[c-'a']++
			} else {
				even[c-'a']--
			}
		}
		m[[2][26]int{odd, even}] = true
	}
	return len(m)
}
