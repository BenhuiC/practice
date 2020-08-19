package partice

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTree(arg []string, n int) *TreeNode {
	if n >= len(arg) || arg[n] == "null" {
		return nil
	}
	r := &TreeNode{}
	i, _ := strconv.ParseInt(arg[n], 10, 64)
	r.Val = int(i)
	r.Left = generateTree(arg, 2*n+1)
	r.Right = generateTree(arg, 2*(n+1))
	return r
}

func pre(root *TreeNode) {
	if root == nil {
		return
	}
	pre(root.Left)
	fmt.Print(root.Val)
	pre(root.Right)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
