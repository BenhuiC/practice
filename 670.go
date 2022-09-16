package partice

import (
	"strconv"
)

func maximumSwap(num int) int {
	res670 := 0
	btNum := []byte(strconv.Itoa(num))
	m := make([]int, 10)
	for i, v := range btNum {
		m[v-'0'] = i
	}
	for i := range btNum {
		for j := 9; j >= 0; j-- {
			if j > int(btNum[i]-'0') && m[j] > i {
				btNum[i], btNum[m[j]] = byte(j+'0'), btNum[i]
				res670, _ = strconv.Atoi(string(btNum))
				return res670
			}
		}
	}
	return num
}
