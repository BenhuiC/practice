package array

import (
	"fmt"
	"testing"
)

func TestConstructor4(t *testing.T) {
	c := Constructor4()
	fmt.Println(c.Remove(0))
	fmt.Println(c.Remove(0))
	fmt.Println(c.Insert(0))
	fmt.Println(c.GetRandom())
	fmt.Println(c.Remove(0))
	fmt.Println(c.Insert(0))
}
