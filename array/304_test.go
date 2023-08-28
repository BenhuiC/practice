package array

import (
	"fmt"
	"testing"
)

func TestConstructor2(t *testing.T) {
	n := Constructor2([][]int{{-1}})
	fmt.Println(n.SumRegion(0, 0, 0, 0))
}
