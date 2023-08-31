package partice

import "sort"

func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)
	if h == len(piles) {
		return piles[h-1]
	}
	l, r := piles[0], piles[h-1]
	for l < r {
		m := (l + r) / 2
		s := 0
		for _, v := range piles {
			s += (v-1)/m + 1
		}
		if s <= h {
			l = m + 1
		} else {
			r = m
		}

	}

	return l
}

type n struct {
	next *n
}
