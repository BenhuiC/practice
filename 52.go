package partice

import (
	"fmt"
	"math"
)

var NQueens int
var QueenMap []int

func totalNQueens(n int) int {
	QueenMap = make([]int, n)
	for i := 0; i < n; i++ {
		QueenMap[i] = -1
	}
	check(0, n)
	return NQueens
}

func check(l, n int) {
	if l >= n {
		NQueens++
		fmt.Println(QueenMap)
		return
	}
	for i := 0; i < n; i++ {
		if !is(l, i, n) {
			continue
		}
		QueenMap[l] = i
		check(l+1, n)
		QueenMap[l] = -1
	}
}

func is(x, y, n int) bool {
	if QueenMap[x] != -1 {
		return false
	}
	for i := 0; i < n; i++ {
		if QueenMap[i] == -1 {
			continue
		}
		if QueenMap[i] == y {
			return false
		}
		if math.Abs(float64(x-i)) == math.Abs(float64(y-QueenMap[i])) {
			return false
		}
	}

	return true
}
