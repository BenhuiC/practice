package offer

import "sort"

func topKFrequent692(words []string, k int) []string {
	sort.Strings(words)
	type pair struct {
		wd  string
		cnt int
	}
	ary := make([]pair, 0, len(words))
	ary = append(ary, pair{words[0], 1})
	for i := 1; i < len(words); i++ {
		if words[i] == words[i-1] {
			ary[len(ary)-1].cnt++
		} else {
			ary = append(ary, pair{words[i], 1})
		}
	}
	sort.Slice(ary, func(i, j int) bool {
		if ary[i].cnt != ary[j].cnt {
			return ary[i].cnt > ary[j].cnt
		}
		return ary[i].wd < ary[j].wd
	})
	res := make([]string, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, ary[i].wd)
	}
	return res
}
