package array

func getSmallestString(s string) string {
	ary := []byte(s)
	for i := 0; i < len(ary)-1; i++ {
		if (ary[i]-'0')&1 == (ary[i+1]-'0')&1 && ary[i] > ary[i+1] {
			ary[i], ary[i+1] = ary[i+1], ary[i]
			break
		}
	}

	return string(ary)
}
