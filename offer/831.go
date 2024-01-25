package offer

import (
	"bytes"
	"strings"
)

func maskPII(s string) string {
	if len(s) == 0 {
		return ""
	}
	maskNum := func(str string) string {
		nums := make([]byte, 0, 10)
		for _, c := range str {
			if '0' <= c && c <= '9' {
				nums = append(nums, byte(c))
			}
		}
		headNu := len(nums) - 10
		res := []byte{}
		if headNu != 0 {
			res = append([]byte{'+'}, bytes.Repeat([]byte{'*'}, headNu)...)
			res = append(res, '-')
		}
		tail := append([]byte("***-***-"), nums[len(nums)-4:]...)
		res = append(res, tail...)
		return string(res)
	}

	maskMail := func(str string) string {
		str = strings.ToLower(str)
		res := make([]byte, 0)
		for i, c := range str {
			if i == 0 {
				res = append(res, byte(c))
				res = append(res, []byte("*****")...)
			} else if str[i+1] == '@' {
				res = append(res, byte(c))
				res = append(res, []byte(str[i+1:])...)
				break
			}
		}
		return string(res)
	}

	if c := s[0]; 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' {
		return maskMail(s)
	} else {
		return maskNum(s)
	}
}
