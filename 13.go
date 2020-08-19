package partice

func romanToInt(s string) int {
	var result int
	vals := map[string]int{"M": 1000, "CM": 900, "D": 500, "CD": 400, "C": 100, "XC": 90, "L": 50, "XL": 40, "X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1}

	for i := 0; i < len(s); {
		j := i + 2
		if j > len(s) {
			j = i + 1
		}
		if v, ok := vals[s[i:j]]; ok {
			result += v
			i = j
			continue
		}
		result += vals[s[i:j-1]]
		i++
	}
	return result
}

func romanToInt2(s string) int {
	var result int
	vals := map[string]int{"M": 1000, "D": 500, "C": 100, "L": 50, "X": 10, "V": 5, "I": 1}

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && vals[s[i:i+1]] < vals[s[i+1:i+2]] {
			result -= vals[s[i:i+1]]
		} else {
			result += vals[s[i:i+1]]
		}
	}
	return result
}
