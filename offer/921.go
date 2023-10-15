package offer

func minAddToMakeValid(s string) int {
	leftNum := 0
	res := 0
	for _, v := range s {
		if v == '(' {
			leftNum++
			continue
		}
		if leftNum > 0 {
			leftNum--
			continue
		}
		res++
	}
	res += leftNum
	return res
}
