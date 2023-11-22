package offer

import (
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	res := 0
	var i, j int
	for i < len(houses) {
		if heaters[j]-res <= houses[i] && heaters[j]+res >= houses[i] {
			i++
			continue
		}
		var tmp int
		if houses[i] < heaters[j] {
			tmp = heaters[j] - houses[i]
			if j > 0 {
				tmp = min(tmp, houses[i]-heaters[j-1])
			}
			i++
		} else {
			tmp = houses[i] - heaters[j]
			if j < len(heaters)-1 {
				tmp = min(tmp, heaters[j+1]-houses[i])
				j++
			}
		}
		if tmp > res {
			res = tmp
		}
	}

	return res
}
