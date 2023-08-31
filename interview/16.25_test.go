package interview

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	l := Constructor(2)
	fmt.Println(l.Get(2))
	l.Put(2, 6)
	fmt.Println(l.Get(1))
	l.Put(1, 5)
	l.Put(1, 2)
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(2))
}
