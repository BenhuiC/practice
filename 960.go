package partice

import (
	"fmt"
)

func minDeletionSize(strs []string) int {
	res := 0
	strLen := len(strs[0])
	dp := make([]int, strLen)
	dp[0] = 1
	for i := 1; i < strLen; i++ {
		for j := 0; j < i; j++ {
			inc := true
			for _, v := range strs {
				if v[i] < v[j] {
					inc = false
					break
				}
			}
			if inc {
				dp[i] = max(dp[i], dp[j]+1)
			} else {
				dp[i] = max(dp[i], 1)
			}
		}
	}
	fmt.Println(dp)
	for _, v := range dp {
		res = max(res, v)
	}

	return strLen - res
}
