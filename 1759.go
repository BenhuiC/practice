package partice

func countHomogenous(s string) int {
	res1759 := 0
	y := 1000000007
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	i, j := 0, 0
	for j < len(s) {
		if s[i] == s[j] {
			res1759 = (res1759 + j - i + 1) % y
			j++
			continue
		}
		i = j
	}

	return res1759
}

/*
a 1
aa 2+1 3
aaa 3+2+1 6
aaaa 4+3+2+1 10
aaaaa 5+4+3+2+1 15

(1+len)*len/2
*/
