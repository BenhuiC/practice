package offer

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	res := make([]string, 0)
	m := make(map[string]int)
	for _, d := range cpdomains {
		a := strings.Split(d, " ")
		cnt, _ := strconv.Atoi(a[0])

		subDomains := strings.Split(a[1], ".")
		for i := range subDomains {
			m[strings.Join(subDomains[i:], ".")] += cnt
		}
	}

	for k, v := range m {
		res = append(res, fmt.Sprintf("%d %s", v, k))
	}

	return res
}
