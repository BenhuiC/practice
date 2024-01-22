package offer

import "fmt"

func ambiguousCoordinates(s string) []string {
	n := len(s)
	s = s[1 : n-1]
	split := func(str string) []string {
		pos := make([]string, 0)
		if str[0] != '0' || str == "0" {
			pos = append(pos, str)
		}
		for p := 1; p < len(str); p++ {
			if p != 1 && str[0] == '0' || str[len(str)-1] == '0' {
				continue
			}
			pos = append(pos, str[:p]+"."+str[p:])
		}
		return pos
	}

	res := make([]string, 0)
	for i := 1; i < n-2; i++ {
		l := split(s[:i])
		r := split(s[i:])
		if len(l) == 0 || len(r) == 0 {
			continue
		}
		for _, x := range l {
			for _, y := range r {
				res = append(res, fmt.Sprintf("(%s, %s)", x, y))
			}
		}
	}
	return res
}
