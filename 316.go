package partice

import "fmt"

// todo
func removeDuplicateLetters(s string) string {
	m := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = i
			continue
		}
		for k := range m {
			if k < s[i] && m[s[i]] < m[k] {
				m[s[i]] = i
				break
			}
		}
	}

	res316 := ""
	for len(m) != 0 {
		t := len(s) + 1
		var r uint8
		for k := range m {
			if m[k] < t {
				r = k
				t = m[k]
			}
		}
		res316 += fmt.Sprintf("%c", r)
		delete(m, r)
	}

	return res316
}
