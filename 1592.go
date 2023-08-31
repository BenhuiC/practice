package partice

import (
	"strings"
)

func reorderSpaces(text string) string {
	res := ""
	ary := strings.Split(text, " ")
	wordLen := 0
	wordCount := 0
	for _, v := range ary {
		if v == "" {
			continue
		}
		ary[wordCount] = v
		wordCount++
		wordLen += len(v)
	}
	ary = ary[:wordCount]
	if len(ary) == 1 {
		return ary[0] + strings.Repeat(" ", len(text)-len(ary[0]))
	}
	in := (len(text) - wordLen) / (wordCount - 1)
	y := (len(text) - wordLen) % (wordCount - 1)
	res = strings.Join(ary, strings.Repeat(" ", in)) + strings.Repeat(" ", y)

	return res
}
