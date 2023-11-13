package offer

import "fmt"

func compress(chars []byte) int {
	n := len(chars)
	idx := 0
	for i := 0; i < n; {
		j := i + 1
		for j < n && chars[j] == chars[i] {
			j++
		}
		chars[idx] = chars[i]
		idx++
		num := fmt.Sprintf("%d", j-i)
		i = j
		if num == "1" {
			continue
		}

		for p := 0; p < len(num); p++ {
			chars[idx] = num[p]
			idx++
		}

	}
	chars = chars[:idx]

	return idx
}
