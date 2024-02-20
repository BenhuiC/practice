package hot

func findAnagrams(s string, p string) []int {
	ls, lp := len(s), len(p)
	if ls < lp {
		return nil
	}
	res := make([]int, 0)
	pAry := [26]int{}
	for _, c := range p {
		pAry[c-'a']++
	}

	l, r := 0, 0
	sAry := [26]int{}

	for r < lp {
		sAry[s[r]-'a']++
		r++
	}
	for r <= ls {
		if sAry == pAry {
			res = append(res, l)
		}
		if r < ls {
			sAry[s[l]-'a']--
			sAry[s[r]-'a']++
		}
		l++
		r++
	}

	return res
}
