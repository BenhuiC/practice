package partice

import (
	"fmt"
	"math"
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {
	if len(arr) <= k {
		return arr
	}
	var res []int
	l, r := 0, 0
	curCg := 0           // 当前k个长度的差值
	minCg := math.MaxInt // 最小差值
	for r < len(arr) {
		if r < k {
			curCg += abs(x - arr[r])
			r++
			continue
		}
		if r == k {
			minCg = curCg
			res = arr[:k]
		}
		curCg += abs(x - arr[r])
		curCg -= abs(x - arr[l])
		l++
		r++
		if minCg > curCg {
			res = arr[l : l+k]
			minCg = curCg
		}
	}

	return res
}

func findClosestElements2(arr []int, k int, x int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return abs(x-arr[i]) <= abs(x-arr[j])
	})
	fmt.Println(arr)
	res := arr[:k]
	sort.Ints(res)
	return res
}
