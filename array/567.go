package array

func checkInclusion(s1 string, s2 string) bool {
	s1len, s2len := len(s1), len(s2)
	if s1len > s2len {
		return false
	}
	count := [26]int{}
	for i, v := range s1 {
		count[v-'a']++
		count[s2[i]-'a']--
	}

	diff := 0
	for _, v := range count {
		if v != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}
	for j, v := range s2[:s2len-s1len] {
		if count[v-'a'] == -1 {
			diff--
		} else if count[v-'a'] == 0 {
			diff++
		}
		count[v-'a']++

		if count[s2[j+s1len]-'a'] == 1 {
			diff--
		} else if count[s2[j+s1len]-'a'] == 0 {
			diff++
		}
		count[s2[j+s1len]-'a']--
		if diff == 0 {
			return true
		}
	}

	return false
}
