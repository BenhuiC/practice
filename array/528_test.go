package array

import (
	"fmt"
	"testing"
)

func TestConstructor3(t *testing.T) {
	s := Constructor3([]int{1, 3})
	fmt.Println(s.PickIndex())
	fmt.Println(s.PickIndex())
	fmt.Println(s.PickIndex())
	fmt.Println(s.PickIndex())
	fmt.Println(s.PickIndex())
	fmt.Println(s.PickIndex())
}
