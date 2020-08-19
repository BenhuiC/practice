package partice

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_AddTwoNumber(t *testing.T) {

	for i := 0; i < 1000; i++ {
		n1 := rand.Intn(1000000)
		n2 := rand.Intn(10000)
		n3 := n1 + n2
		l1 := generatorList(n1)
		l2 := generatorList(n2)
		l3 := addTwoNumbers(l1, l2)
		if !equal(l3, generatorList(n3)) {
			t.Error(fmt.Sprintf("index:%d\twant:%d+%d=%d\tget:%d", i, n1, n2, n3, valueOf(l3)))
		}
	}
}

func generatorList(n int) *ListNode {
	l := &ListNode{}
	if n < 10 {
		l.Val = n
		return l
	}
	t := n % 10
	l.Val = t
	l.Next = generatorList(n / 10)
	return l
}

func valueOf(l *ListNode) int {
	if l == nil {
		return 0
	}
	return l.Val + 10*valueOf(l.Next)
}

func equal(l1, l2 *ListNode) bool {
	if l1 == l2 {
		return true
	}
	if l1.Val == l2.Val {
		return equal(l1.Next, l2.Next)
	}
	return false
}
