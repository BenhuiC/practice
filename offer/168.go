package offer

import "fmt"

func convertToTitle(columnNumber int) string {
	res := ""
	for columnNumber != 0 {
		x := (columnNumber - 1) % 26
		res = fmt.Sprintf("%c%s", 'A'+x, res)
		columnNumber /= 26
		if x == 25 {
			columnNumber--
		}
	}
	return res
}
