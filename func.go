package partice

import "fmt"

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(m, n int) int {
	if m < n {
		return m
	}
	return n
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func AbsFloat64(f float64) float64 {
	if f < 0 {
		return -f
	}
	return f
}

func printMatrix(m [][]int) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Printf("%d\t", m[i][j])
		}
		fmt.Println()
	}
}
