package partice

import "fmt"

func multiply(num1 string, num2 string) string {
	var res string
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	if num1 == "1" {
		return num2
	}
	if num2 == "1" {
		return num1
	}

	tmps := make([]string, 0)
	zero := ""
	for i := len(num2) - 1; i >= 0; i-- {
		t := singleStrMult(num1, num2[i]-'0')
		if t != "" {
			tmps = append(tmps, t+zero)
		}
		zero += "0"
	}
	for i := 0; i < len(tmps); i++ {
		res = strAdd(res, tmps[i])
	}

	return res
}

func singleStrMult(num1 string, v uint8) string {
	res := ""
	if v == 0 {
		return ""
	}
	if v == 1 {
		return num1
	}
	var y, c uint8
	for i := len(num1) - 1; i >= 0; i-- {
		tmp := v*(num1[i]-'0') + y
		y = tmp / 10
		c = tmp % 10
		res = fmt.Sprintf("%c", c+'0') + res
	}
	if y != 0 {
		res = fmt.Sprintf("%c", y+'0') + res
	}

	return res
}

func strAdd(num1, num2 string) string {
	if num1 == "" || num1 == "0" {
		return num2
	}
	if num2 == "" || num2 == "0" {
		return num1
	}
	var res string
	var l, s string
	if len(num1) > len(num2) {
		l = num1
		s = num2
	} else {
		l = num2
		s = num1
	}
	lo := len(l) - 1
	so := len(s) - 1
	var y, c uint8
	for lo >= 0 {
		var tmp uint8
		if so >= 0 {
			tmp = l[lo] - '0' + s[so] - '0' + y
		} else {
			tmp = l[lo] - '0' + y
		}
		c = tmp % 10
		y = tmp / 10
		res = fmt.Sprintf("%c", c+'0') + res
		lo--
		so--
	}
	if y != 0 {
		res = fmt.Sprintf("%c", y+'0') + res
	}

	return res
}
