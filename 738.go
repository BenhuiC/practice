package partice

import (
	"fmt"
	"math"
)

func monotoneIncreasingDigits(n int) int {
	res738 := 0
	strN := fmt.Sprint(n)
	nlen := len(strN)
	i := 0
	for i+1 < nlen && strN[i+1] >= strN[i] {
		i++
	}
	if i == nlen-1 {
		return n
	}
	for i-1 >= 0 && strN[i-1] == strN[i] {
		i--
	}
	p := nlen - i - 1
	pw := int(math.Pow10(p))
	res738 = (n/pw-1)*pw + pw - 1
	return res738
}
