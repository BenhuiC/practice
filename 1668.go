package partice

import "strings"

func maxRepeating(sequence string, word string) int {
	res := 0
	for {
		if strings.Contains(sequence, strings.Repeat(word, res+1)) {
			res++
		} else {
			break
		}
	}

	return res
}
