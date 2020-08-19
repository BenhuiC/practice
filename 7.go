package partice

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {
	var result int
	tp := fmt.Sprint(x)
	strVal := ""
	if tp[0] == '-' {
		strVal = "-" + reverseStr(tp[1:])
	} else {
		strVal = reverseStr(tp)
	}
	val, _ := strconv.ParseInt(strVal, 10, 64)
	if val > math.MaxInt32 || val < math.MinInt32 {
		result = 0
	} else {
		result = int(val)
	}
	return result
}

func reverse2(x int) int {
	var result int64
	for x != 0 {
		result = result*10 + int64(x%10)
		x = x / 10
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return int(result)
}

func reverseStr(s string) (result string) {
	if len(s) <= 1 {
		return s
	}
	re := make([]string, len(s))
	i, j := 0, len(s)-1
	for i <= j {
		re[j] = fmt.Sprintf("%c", s[i])
		re[i] = fmt.Sprintf("%c", s[j])
		i++
		j--
	}
	result = strings.Join(re, "")
	return
}
