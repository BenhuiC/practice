package partice

import (
	"fmt"
	"math"
	"strconv"
)

func myAtoi(str string) int {
	var start int
	var tp string
	for _, s := range str {
		if s == ' ' {
			start++
		} else {
			break
		}
	}
	if start >= len(str) {
		return 0
	}
	if str[start] == '-' || str[start] == '+' {
		tp += fmt.Sprintf("%c", str[start])
		start++
	} else if str[start] < 48 || str[start] > 57 {
		return 0
	}
	for start < len(str) && str[start] >= 48 && str[start] <= 57 {
		tp = tp + fmt.Sprintf("%c", str[start])
		start++
	}
	val, _ := strconv.ParseInt(tp, 10, 64)
	if val > math.MaxInt32 {
		return math.MaxInt32
	}
	if val < math.MinInt32 {
		return math.MinInt32
	}
	return int(val)
}
