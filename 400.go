package partice

import (
	"math"
)

func findNthDigit(n int) int {
	d := 1
	for count := 9; n > d*count; count *= 10 {
		n -= d * count
		d++
	}
	index := n - 1
	start := int(math.Pow10(d - 1))
	num := start + index/d
	digitIndex := index % d
	return num / int(math.Pow10(d-digitIndex-1)) % 10
}
