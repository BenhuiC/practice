package types

import (
	"fmt"
	"testing"
)

func TestConstructorMapSum(t *testing.T) {
	c := ConstructorMapSum()
	c.Insert("apple", 3)
	fmt.Println(c.Sum("ap"))
	c.Insert("app", 2)
	fmt.Println(c.Sum("ap"))
}
