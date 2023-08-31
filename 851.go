package partice

var (
	mp851  map[int][]int
	res851 []int
	q      []int
)

func loudAndRich(richer [][]int, quiet []int) []int {
	mp851 = make(map[int][]int)
	for _, v := range richer {
		if mp851[v[1]] == nil {
			mp851[v[1]] = make([]int, 0)
		}
		mp851[v[1]] = append(mp851[v[1]], v[0])
	}
	res851 = make([]int, len(quiet))
	q = quiet
	for i := range res851 {
		res851[i] = -1
	}
	for i := range res851 {
		if res851[i] != -1 {
			continue
		}
		res851[i] = quite(i)
	}

	return res851
}

func quite(i int) int {
	if res851[i] != -1 {
		return res851[i]
	}
	res := i
	for _, v := range mp851[i] {
		if tmp := quite(v); q[tmp] < q[res] {
			res = tmp
		}
	}
	res851[i] = res
	return res
}
