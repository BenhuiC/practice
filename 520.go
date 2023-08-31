package partice

func detectCapitalUse(word string) bool {
	if word == "" {
		return false
	}
	firstCap := word[0] >= 'A' && word[0] <= 'Z'
	if !firstCap {
		for i := 1; i < len(word); i++ {
			if word[0] >= 'A' && word[0] <= 'Z' {
				return false
			}
		}
		return true
	}
	for i := 2; i < len(word); i++ {
		n := int(word[i]) - int(word[i-1])
		if n >= 26 || n <= -26 {
			return false
		}
	}
	return true
}
