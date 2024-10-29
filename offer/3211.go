package offer

func validStrings(n int) []string {
	cur := []string{"0", "1"}
	for i := 1; i < n; i++ {
		next := make([]string, 0)
		for _, c := range cur {
			if c[len(c)-1] == '0' {
				next = append(next, c+"1")
			} else {
				next = append(next, c+"0", c+"1")
			}
		}
		cur = next
	}
	return cur
}
