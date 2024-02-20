package hot

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	res := make([]string, 0)
	m := map[uint8][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	var bk func(s string, idx int)
	bk = func(s string, idx int) {
		if idx >= len(digits) {
			res = append(res, s)
			return
		}
		for _, c := range m[digits[idx]] {
			bk(s+c, idx+1)
		}
	}
	bk("", 0)

	return res
}
