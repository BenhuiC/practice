package array

import (
	"fmt"
	"testing"
)

func TestConstructor5(t *testing.T) {
	s := Constructor5(7, []int{2, 3, 5})
	fmt.Println(s.Pick())
	fmt.Println(s.Pick())
	fmt.Println(s.Pick())
	fmt.Println(s.Pick())
}
