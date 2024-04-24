package array

import (
	"sort"
)

func deckRevealedIncreasing(deck []int) []int {
	var genIndex func(n int, odd bool) []int
	genIndex = func(n int, odd bool) []int {
		if n == 1 {
			return []int{1}
		}
		idxs := make([]int, 0, n)
		var next []int
		if n&1 == 1 {
			if odd {
				next = genIndex(n/2, false)
			} else {
				next = genIndex((n+1)/2, true)
			}
		} else {
			next = genIndex(n/2, odd)
		}
		if odd {
			for i := 1; i <= n; i += 2 {
				idxs = append(idxs, i)
			}
			for _, nidx := range next {
				idxs = append(idxs, 2*nidx)
			}
		} else {
			for i := 2; i <= n; i += 2 {
				idxs = append(idxs, i)
			}
			for _, nidx := range next {
				idxs = append(idxs, 2*nidx-1)
			}
		}

		return idxs
	}

	idxs := genIndex(len(deck), true)
	sort.Ints(deck)
	res := make([]int, len(deck))
	for i := range deck {
		res[idxs[i]-1] = deck[i]
	}
	return res
}

func deckRevealedIncreasing2(deck []int) []int {
	n := len(deck)
	res := make([]int, n)
	idxs := make([]int, 0, n)
	for i := 0; i < n; i++ {
		idxs = append(idxs, i)
	}
	sort.Ints(deck)

	for _, d := range deck {
		res[idxs[0]] = d
		idxs = idxs[1:]
		if len(idxs) > 0 {
			v := idxs[0]
			idxs = append(idxs[1:], v)
		}
	}

	return res
}
