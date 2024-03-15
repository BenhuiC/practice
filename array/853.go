package array

import "sort"

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	if n < 2 {
		return 1
	}
	type item struct {
		idx int
		pos int
	}
	ary := make([]item, 0, n)
	for i, p := range position {
		ary = append(ary, item{idx: i, pos: p})
	}
	sort.Slice(ary, func(i, j int) bool {
		return ary[i].pos < ary[j].pos
	})

	res := 1
	r := n - 1
	for r > 0 {
		rTime := float64(target-ary[r].pos) / float64(speed[ary[r].idx])
		l := r - 1
		for l >= 0 {
			lTime := float64(target-ary[l].pos) / float64(speed[ary[l].idx])
			if lTime <= rTime {
				l--
			} else {
				res++
				break
			}
		}
		r = l
	}

	return res
}
