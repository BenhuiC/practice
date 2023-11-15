package offer

import (
	"bytes"
	"sort"
)

func frequencySort(s string) string {
	m := make(map[byte]int)
	for _, v := range s {
		m[byte(v)]++
	}

	type pair struct {
		c byte
		n int
	}
	pairs := make([]pair, 0, len(m))
	for k, v := range m {
		pairs = append(pairs, pair{c: k, n: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].n > pairs[j].n
	})

	res := make([]byte, 0, len(s))
	for _, p := range pairs {
		res = append(res, bytes.Repeat([]byte{p.c}, p.n)...)
	}
	return string(res)
}
