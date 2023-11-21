package offer

import (
	"strconv"
	"strings"
)

func validIPAddress(queryIP string) string {
	if isIPv4(queryIP) {
		return "IPv4"
	} else if isIPv6(queryIP) {
		return "IPv6"
	}
	return "Neither"
}

func isIPv4(queryIP string) bool {
	nuAry := strings.Split(queryIP, ".")
	if len(nuAry) != 4 {
		return false
	}
	for _, n := range nuAry {
		if len(n) == 0 || len(n) > 3 {
			return false
		}
		if len(n) > 1 && n[0] == '0' {
			return false
		}
		nv, err := strconv.Atoi(n)
		if err != nil || nv > 255 {
			return false
		}
	}
	return true
}

func isIPv6(queryIP string) bool {
	nuAry := strings.Split(queryIP, ":")
	if len(nuAry) != 8 {
		return false
	}
	for _, n := range nuAry {
		if len(n) == 0 || len(n) > 4 {
			return false
		}
		if _, err := strconv.ParseUint(n, 16, 64); err != nil {
			return false
		}
	}

	return true
}
