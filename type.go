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

func (l *ListNode) Print() {
	for t := l; t != nil; t = t.Next {
		fmt.Printf("val:%d\t", t.Val)
	}
	fmt.Println()
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

func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}

type MinHeap []*ListNode

func NewMinHeap() MinHeap {
	m := make([]*ListNode, 0)
	return m
}

func (m *MinHeap) Sink(n int) {
	for n*2+1 < len(*m) {
		if (*m)[n*2+1].Val < (*m)[n].Val || n*2+2 < len(*m) && (*m)[n*2+2].Val < (*m)[n].Val {
			if n*2+2 < len(*m) && (*m)[n*2+2].Val < (*m)[n*2+1].Val {
				(*m)[n*2+2], (*m)[n] = (*m)[n], (*m)[n*2+2]
				n = n*2 + 2
			} else {
				(*m)[n*2+1], (*m)[n] = (*m)[n], (*m)[n*2+1]
				n = n*2 + 1
			}
		} else {
			break
		}
	}
}

func (m *MinHeap) Up(n int) {
	for (*m)[n].Val < (*m)[(n-1)/2].Val && n >= 0 {
		(*m)[n], (*m)[(n-1)/2] = (*m)[(n-1)/2], (*m)[n]
		n = (n - 1) / 2
	}
}

func (m *MinHeap) Insert(node *ListNode) {
	*m = append(*m, node)
	m.Up(len(*m) - 1)
}

func (m *MinHeap) DelMin() {
	l := len(*m)
	(*m)[0] = (*m)[l-1]
	*m = (*m)[:l-1]
	m.Sink(0)
}

func (m *MinHeap) Print() {
	for i := 0; i < len(*m); i++ {
		fmt.Printf("%d\t", (*m)[i].Val)
	}
	fmt.Println()
}
