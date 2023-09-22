package back_trace

import "fmt"

func printMatrix(ary [][]int) {
	m := len(ary)
	if m == 0 {
		return
	}
	n := len(ary[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d\t", ary[i][j])
		}
		fmt.Println()
	}
}
