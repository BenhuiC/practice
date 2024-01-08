package offer

import "math/bits"

func countPrimeSetBits(left, right int) (ans int) {
	for x := left; x <= right; x++ {
		if 1<<bits.OnesCount(uint(x))&665772 != 0 {
			ans++
		}
	}
	return
}
