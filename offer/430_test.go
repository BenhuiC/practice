package offer

import (
	"fmt"
	"testing"
)

func Test_flatten(t *testing.T) {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}
	node7 := &Node{Val: 7}
	node8 := &Node{Val: 8}
	node9 := &Node{Val: 9}
	node10 := &Node{Val: 10}
	node11 := &Node{Val: 11}
	node12 := &Node{Val: 12}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node2.Prev = node1
	node3.Prev = node2
	node4.Prev = node3
	node5.Prev = node4
	node6.Prev = node5

	node7.Next = node8
	node8.Next = node9
	node9.Next = node10
	node8.Prev = node7
	node9.Prev = node8
	node10.Prev = node9

	node11.Next = node12
	node12.Prev = node11

	node3.Child = node7
	node8.Child = node11

	got := flatten(node1)
	for n := got; n != nil; n = n.Next {
		fmt.Printf("pre %v\tcur %v\tnext %v\n", n.Prev, n.Val, n.Next)
	}
}
