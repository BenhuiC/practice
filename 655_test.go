package partice

import (
	"fmt"
	"test/practice"
	"testing"
)

func Test_printTree(t *testing.T) {
	tests := []struct {
		root *practice.TreeNode
	}{
		{
			root: &practice.TreeNode{Val: 1},
		},
		{
			root: &practice.TreeNode{Val: 1, Left: &practice.TreeNode{Val: 2}, Right: &practice.TreeNode{Val: 3}},
		},
		{
			root: &practice.TreeNode{Val: 1, Left: &practice.TreeNode{Val: 2}},
		},
		{
			root: &practice.TreeNode{Val: 1, Left: &practice.TreeNode{Val: 2, Left: &practice.TreeNode{Val: 3}}},
		},
	}
	for _, tt := range tests {
		got := printTree(tt.root)
		fmt.Println()
		pintMat(got)
	}
}

func pintMat(g [][]string) {
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[0]); j++ {
			if g[i][j] == "" {
				fmt.Print(" N ")
			} else {
				fmt.Printf(" %s ", g[i][j])
			}
		}
		fmt.Println()
	}
}
