package array

import (
	"slices"
)

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	bitMap := make(map[int]bool)
	i := 0
	for 1<<(i+1) <= n {
		i++
	}
	maxBit := i
	for v := n; v >= 0 && i >= 0; {
		if v >= 1<<i {
			bitMap[i] = true
			v -= 1 << i
		}
		i--
	}

	res := make([]byte, 0, i+1)
	for idx := 0; idx <= maxBit; {
		if bitMap[idx] {
			res = append(res, '1')
			if idx&1 == 0 {
				idx++
				continue
			}

			tmp := idx + 1
			for bitMap[tmp] {
				tmp += 1
				res = append(res, '0')
			}
			bitMap[tmp] = true
			if tmp > maxBit {
				maxBit = tmp
			}
			idx = tmp
		} else {
			res = append(res, '0')
			idx++
		}
	}

	slices.Reverse(res)
	return string(res)
}
