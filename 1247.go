package partice

func minimumSwap(s1 string, s2 string) int {
	if len(s1) != len(s2) {
		return -1
	}
	if s1 == s2 {
		return 0
	}
	res1247 := 0
	ary1 := make([]uint8, len(s1))
	ary2 := make([]uint8, len(s2))
	for i := 0; i < len(s1); i++ {
		ary1[i] = s1[i]
		ary2[i] = s2[i]
	}

	i := 0
	for {
		if i >= len(s1) {
			break
		}
		if ary1[i] == ary2[i] {
			i++
			continue
		}
		end := true
		best := 0
		for j := i + 1; j < len(s1); j++ {
			if ary1[j] == ary2[j] {
				continue
			}
			end = false
			if ary2[j] == ary2[i] {
				best = j
			}
		}
		if end {
			return -1
		}
		if best != 0 {
			ary2[best] = ary1[i]
			ary1[i] = ary2[i]
		} else {
			ary1[i], ary2[i] = ary2[i], ary1[i]
		}
		res1247++
	}

	return res1247
}
