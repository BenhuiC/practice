package _16

import "fmt"

type Node struct {
    Val int
    Left *Node
    Right *Node
    Next *Node
}

var t *Node

func connect(root *Node) *Node {
	if root==nil{
		return root
	}
	connect(root.Left)
	connect(root.Right)
	if root.Left==nil&&root.Right==nil{
		if t!=nil{
			t.Next=root
			t=root
		}else{
			t=root
		}
	}
	fmt.Println(root.Val)
	return root
}

