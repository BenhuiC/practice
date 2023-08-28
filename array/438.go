package array

func findAnagrams(s string, p string) []int {
	slen, plen := len(s), len(p)
	if slen < plen {
		return nil
	}
	count := [26]int{}
	for i, ch := range p {
		count[ch-'a']--
		count[s[i]-'a']++
	}
	diff := 0
	for _, v := range count {
		if v != 0 {
			diff++
		}
	}
	res := make([]int, 0)
	if diff == 0 {
		res = append(res, 0)
	}

	for i, v := range s[:slen-plen] {
		if count[v-'a'] == 1 {
			diff--
		} else if count[v-'a'] == 0 {
			diff++
		}
		count[v-'a']--

		if count[s[i+plen]-'a'] == -1 {
			diff--
		} else if count[s[i+plen]-'a'] == 0 {
			diff++
		}
		count[s[i+plen]-'a']++
		if diff == 0 {
			res = append(res, i+1)
		}
	}

	return res
}
