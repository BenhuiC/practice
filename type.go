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

type stack []int

func (s *stack) Pop() int {
	if len(*s) > 0 {
		r := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return r
	}
	return 0
}

func (s *stack) Peek() int {
	if len(*s) > 0 {
		return (*s)[len(*s)-1]
	}
	return 0
}

func (s *stack) Empty() bool {
	return s == nil || len(*s) <= 0
}
