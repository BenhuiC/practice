package offer

func expressiveWords(s string, words []string) int {
	n := len(words)
	if n == 0 {
		return 0
	}

	canExp := func(src, target string) bool {
		if src == target {
			return true
		}
		l1, l2 := len(src), len(target)
		i, j := 0, 0
		for i < l1 && j < l2 {
			if src[i] != target[j] {
				return false
			}
			il := 0
			for ir := i; ir < l1 && src[ir] == src[i]; ir++ {
				il++
			}
			jl := 0
			for jr := j; jr < l2 && target[jr] == target[j]; jr++ {
				jl++
			}
			if il != jl && jl < 3 || jl < il {
				return false
			}
			i += il
			j += jl
		}
		if i < l1 || j < l2 {
			return false
		}
		return true
	}

	res := 0
	for _, w := range words {
		if canExp(w, s) {
			res++
		}
	}
	return res
}
