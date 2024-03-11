package back_trace

import (
	"fmt"
	"math"
)

func splitIntoFibonacci(num string) []int {
	n := len(num)
	if n < 3 {
		return nil
	}
	res := make([]int, 0)
	var bk func(idx int) bool
	bk = func(idx int) bool {
		if idx >= n {
			if len(res) >= 3 {
				return true
			}
			return false
		}
		rl := len(res)
		if num[idx] == '0' {
			if rl < 2 || rl >= 2 && res[rl-2]+res[rl-1] == 0 {
				res = append(res, 0)
				if bk(idx + 1) {
					return true
				}
				res = res[:len(res)-1]
			}
			return false
		}

		val := 0
		for i := idx; i < n; i++ {
			val = val*10 + int(num[i]-'0')
			if val > math.MaxInt32 {
				return false
			}
			if rl < 2 {
				res = append(res, val)
				if bk(i + 1) {
					return true
				}
				res = res[:len(res)-1]
			} else {
				if res[rl-2]+res[rl-1] > val {
					continue
				}
				if res[rl-2]+res[rl-1] < val {
					return false
				}

				res = append(res, val)
				fmt.Println("--------", res)
				if bk(i + 1) {
					return true
				}
				res = res[:len(res)-1]
			}
		}
		return false
	}

	if bk(0) {
		return res
	}
	return nil
}
