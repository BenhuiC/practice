package partice

import "sort"

func makesquare(matchsticks []int) bool {
	all := 0
	for _, v := range matchsticks {
		all += v
	}
	if all%4 != 0 {
		return false
	}
	sort.Ints(matchsticks)
	avg := all / 4
	if matchsticks[len(matchsticks)-1] > avg {
		return false
	}
	n := len(matchsticks)
	st := make([]int, 1<<n)
	mask := 1<<n - 1
	var dfs func(s, v int) bool
	dfs = func(s, v int) bool {
		if s == mask {
			return true
		}
		if st[s] != 0 {
			return st[s] == 1
		}
		for i, num := range matchsticks {
			if s>>i&1 == 1 {
				continue
			}
			if v+num > avg {
				break
			}
			if dfs(s|1<<i, (v+num)%avg) {
				st[s] = 1
				return true
			}
		}
		st[s] = -1
		return false
	}
	return dfs(0, 0)
}
