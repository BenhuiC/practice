package partice

import "fmt"

func isPalindrome(x int) bool {
	str := fmt.Sprintf("%d", x)
	i, j := 0, len(str)-1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
