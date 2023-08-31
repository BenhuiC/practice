package partice

import "sort"

func carPooling(trips [][]int, capacity int) bool {
	sort.Slice(trips, func(i, j int) bool {
		return trips[i][1] < trips[j][1]
	})
	last := capacity
	mp := make(map[int]int)
	for _, v := range trips {
		for k := range mp {
			if k <= v[1] {
				last += mp[k]
				delete(mp, k)
			}
		}
		if v[0] > last {
			return false
		}
		last -= v[0]
		mp[v[2]] = mp[v[2]] + v[0]
	}

	return true
}
