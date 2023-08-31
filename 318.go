package partice

func maxProduct(words []string) int {
	ary := make([]int, 0)
	for i := 0; i < len(words); i++ {
		nu := 0
		for _, s := range words[i] {
			nu = nu | 2<<(s-'a'+1)
		}
		ary = append(ary, nu)
	}
	res318 := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if ary[i]&ary[j] == 0 {
				res318 = max(res318, len(words[i])*len(words[j]))
			}
		}
	}
	return res318
}
