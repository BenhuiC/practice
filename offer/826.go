package offer

import "sort"

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	n := len(difficulty)
	type item struct {
		diff, prof int
	}
	profs := make([]item, 0, n)
	for i := 0; i < n; i++ {
		profs = append(profs, item{difficulty[i], profit[i]})
	}
	sort.Slice(profs, func(i, j int) bool {
		if profs[i].prof == profs[j].prof {
			return profs[i].diff < profs[j].diff
		} else {
			return profs[i].prof < profs[j].prof
		}
	})

	m := len(worker)
	sort.Ints(worker)
	res := 0
	idx := n - 1
	for i := m - 1; i >= 0; i-- {
		for idx >= 0 && profs[idx].diff > worker[i] {
			idx--
		}
		if idx >= 0 {
			res += profs[idx].prof
		}
	}

	return res
}
