package partice

import "fmt"

func convert(s string, numRows int) string {
	result := ""
	if len(s) <= numRows {
		return s
	}
	if numRows <= 1 {
		return s
	}

	arg := make([]string, numRows, numRows)
	for i := 0; i < len(s); i++ {
		ans := i / (numRows - 1) % 2
		cur := i % (numRows - 1)
		if ans == 0 {
			arg[cur] += fmt.Sprintf("%c", s[i])
		} else {
			arg[numRows-1-cur] += fmt.Sprintf("%c", s[i])
		}
	}
	for _, j := range arg {
		result += j
	}

	return result
}
