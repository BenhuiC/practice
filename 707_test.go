package partice

import (
	"fmt"
	"testing"
)

func TestConstructor707(t *testing.T) {
	l := Constructor707()
	l.AddAtTail(1)
	fmt.Println(l.Get(0))

	l.AddAtHead(7)
	l.AddAtHead(2)
	l.AddAtHead(1)
	l.AddAtIndex(3, 0)
	l.DeleteAtIndex(2)
	l.AddAtHead(6)
	l.AddAtTail(4)
	fmt.Println(l.Get(4))
	l.AddAtHead(4)
	l.AddAtIndex(5, 0)
	l.AddAtHead(6)

	l.AddAtIndex(0, 10)
	l.AddAtIndex(0, 20)
	l.AddAtIndex(1, 30)
	fmt.Println(l.Get(0))
}
