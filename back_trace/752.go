package back_trace

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	deadMap := make(map[string]bool)
	for _, v := range deadends {
		if v == target {
			return -1
		}
		if v == "0000" {
			return -1
		}
		deadMap[v] = true
	}

	seen := make(map[string]bool)
	get := func(stat string) (ret []string) {
		bts := []byte(stat)
		for i, b := range bts {
			val := b - 1
			if val < '0' {
				val = '9'
			}
			bts[i] = val
			ret = append(ret, string(bts))
			val = b + 1
			if val > '9' {
				val = '0'
			}
			bts[i] = val
			ret = append(ret, string(bts))
			bts[i] = b
		}
		return
	}

	seen["0000"] = true
	type pair struct {
		stat string
		step int
	}
	ary := []pair{{stat: "0000", step: 0}}
	for len(ary) != 0 {
		p := ary[0]
		ary = ary[1:]
		for _, v := range get(p.stat) {
			if seen[v] || deadMap[v] {
				continue
			}
			if v == target {
				return p.step + 1
			}
			seen[v] = true
			ary = append(ary, pair{stat: v, step: p.step + 1})
		}
	}

	return -1
}
