package partice

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	str := countAndSay(n - 1)
	result := ""
	for i := 0; i < len(str); {
		var j = i
		for ; j < len(str) && str[j] == str[i]; j++ {
		}
		result += fmt.Sprintf("%d%d", j-i, str[i]-48)
		i = j
	}
	return result
}
