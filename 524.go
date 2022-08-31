package partice

func findLongestWord(s string, dictionary []string) string {
	var res string
	m := make(map[uint8][]int)
	for i := range s {
		_, ok := m[s[i]]
		if !ok {
			m[s[i]] = make([]int, 0)
		}
		m[s[i]] = append(m[s[i]], i)
	}
	for _, v := range dictionary {
		if len(res) > len(v) {
			continue
		}
		cur := -1
		has := false
		for t := range v {
			has = false
			ary, ok := m[v[t]]
			if !ok || len(ary) == 0 {
				break
			}
			for _, val := range ary {
				if val > cur {
					cur = val
					has = true
					break
				}
			}
			if !has {
				break
			}
		}
		if has {
			if len(v) > len(res) {
				res = v
			} else if len(v) == len(res) && v < res {
				res = v
			}
		}
	}

	return res
}
