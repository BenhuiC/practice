package offer

import "bytes"

func customSortString(order string, s string) string {
	m := make(map[byte]int)
	for _, v := range s {
		m[byte(v)]++
	}
	res := make([]byte, 0, len(s))
	for _, c := range order {
		cnt := m[byte(c)]
		if cnt == 0 {
			continue
		}
		res = append(res, bytes.Repeat([]byte{byte(c)}, cnt)...)
		delete(m, byte(c))
	}
	for k, v := range m {
		if v > 0 {
			res = append(res, bytes.Repeat([]byte{k}, v)...)
		}
	}

	return string(res)
}
