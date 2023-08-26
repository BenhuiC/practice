package link

import (
	"fmt"
	"testing"
)

func Test_minheap_push(t *testing.T) {
	m := minheap{}
	for i := 10; i > 0; i-- {
		m.push(&ListNode{Val: i})
	}
	for {
		v := m.pop()
		if v == nil {
			break
		}
		fmt.Println(v.Val)
	}
}
