package offer

func titleToNumber(columnTitle string) int {
	w := 1
	res := 0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		res += int(columnTitle[i]-'A'+1) * w
		w *= 26
	}
	return res
}
