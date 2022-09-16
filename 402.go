package partice

import (
	"fmt"
	"strings"
)

func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}
	res402 := ""
	stack := make([]uint8, 0)
	for i := 0; i < len(num); i++ {
		for len(stack) != 0 && k > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	for _, v := range stack {
		res402 += fmt.Sprintf("%c", v)
	}
	res402 = res402[:len(res402)-k]
	res402 = strings.TrimLeft(res402, "0")
	if len(res402) == 0 {
		return "0"
	}
	return res402
}
