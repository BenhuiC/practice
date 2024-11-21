package offer150

func intToRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	char := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res := ""
	idx := 0
	for num > 0 && idx < len(val) {
		if num >= val[idx] {
			res += char[idx]
			num -= val[idx]
		} else {
			idx++
		}
	}

	return res
}
