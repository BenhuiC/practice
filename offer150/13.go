package offer150

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0
	for i, c := range s {
		if i != len(s)-1 && m[s[i+1]] > m[byte(c)] {
			res -= m[byte(c)]
		} else {
			res += m[byte(c)]
		}
	}
	return res
}
