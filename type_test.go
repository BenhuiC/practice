package partice

import (
	"fmt"
	"testing"
)

func Test_generateTree(t *testing.T) {
	type args struct {
		arg []string
		n   int
	}
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "1",
			args: []string{"-1", "0", "3", "-2", "4", "null", "null", "8"},
		},
	}
	for _, tt := range tests {
		fmt.Println("----------------")
		pre(generateTree(tt.args, 0))
		fmt.Println("\n----------------")
	}
}

func TestNewMinHeap(t *testing.T) {
	m := NewMinHeap()
	m.Insert(&ListNode{Val: 1})
	m.Insert(&ListNode{Val: 4})
	m.Insert(&ListNode{Val: 20})
	m.Insert(&ListNode{Val: 3})
	m.Insert(&ListNode{Val: 7})
	m.Insert(&ListNode{Val: 2})
	m.Insert(&ListNode{Val: 12})
	m.Insert(&ListNode{Val: 13})
	m.Insert(&ListNode{Val: 14})
	m.Insert(&ListNode{Val: 15})
	m.Print()
	m.DelMin()
	m.Print()
}
