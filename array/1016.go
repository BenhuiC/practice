package array

import (
	"strconv"
	"strings"
)

func queryString(s string, n int) bool {
	var i int64
	for i = 1; i <= int64(n); i++ {
		if !strings.Contains(s, strconv.FormatInt(i, 2)) {
			return false
		}
	}
	return true
}
