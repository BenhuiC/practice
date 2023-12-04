package offer

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	ary1, ary2 := strings.Split(num1, "+"), strings.Split(num2, "+")
	s1, _ := strconv.Atoi(ary1[0])
	x1, _ := strconv.Atoi(ary1[1][:len(ary1[1])-1])
	s2, _ := strconv.Atoi(ary2[0])
	x2, _ := strconv.Atoi(ary2[1][:len(ary2[1])-1])

	s := s1*s2 - x1*x2
	x := s1*x2 + x1*s2
	res := fmt.Sprintf("%d+%di", s, x)
	return res
}
