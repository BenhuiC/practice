package utils

import "fmt"

func PrintMatrix(m [][]int) {
	for _, row := range m {
		for _, v := range row {
			fmt.Printf("%d\t", v)
		}
		fmt.Println()
	}
}
