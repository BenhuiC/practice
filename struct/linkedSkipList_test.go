package _struct

import (
	"fmt"
	"testing"
)

func TestNewLinkedSkipList(t *testing.T) {
	s := NewLinkedSkipList[int](4)
	fmt.Println(s)
}
