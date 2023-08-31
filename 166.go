package partice

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	var res166 string
	if (numerator ^ denominator) < 0 {
		res166 += "-"
		if numerator < 0 {
			numerator = -numerator
		}
		if denominator < 0 {
			denominator = -denominator
		}
	}
	// 整数位
	res166 += fmt.Sprint(numerator / denominator)
	numerator = (numerator % denominator) * 10
	if numerator == 0 {
		return res166
	}
	res166 += "."
	m := map[int]int{}
	for numerator != 0 {
		if v, ok := m[numerator]; ok {
			res166 = res166[:v] + "(" + res166[v:] + ")"
			break
		}
		m[numerator] = len(res166)
		num := numerator / denominator
		numerator = (numerator % denominator) * 10
		res166 += strconv.Itoa(num)
	}
	return res166
}
