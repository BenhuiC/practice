package partice

import (
	"fmt"
	"testing"
)

func TestTreeAncestor_GetKthAncestor(t *testing.T) {
	ary := []int{}
	a := ConstructorTreeAncestor(len(ary), ary)
	fmt.Println(a.GetKthAncestor(3, 1))
	fmt.Println(a.GetKthAncestor(5, 2))
	fmt.Println(a.GetKthAncestor(6, 3))
}
