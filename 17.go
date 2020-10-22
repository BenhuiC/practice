package partice

var letterResult []string
var m = map[uint8][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	letterResult = make([]string, 0)
	if len(digits) == 0 {
		return letterResult
	}
	letter(digits, "", 0)
	return letterResult
}

func letter(digits, str string, index int) {
	if index == len(digits) {
		letterResult = append(letterResult, str)
		return
	}
	t := m[digits[index]]
	for _, v := range t {
		str += v
		letter(digits, str, index+1)
		str = str[:len(str)-1]
	}
}
