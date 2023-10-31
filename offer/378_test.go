package offer

import (
	"fmt"
	"testing"
)

func Test_kthSmallest(t *testing.T) {
	fmt.Println(kthSmallest([][]int{{1, 5, 9, 10, 12}, {10, 11, 13, 20, 20}, {12, 13, 15, 20, 20}, {12, 13, 15, 20, 20}, {12, 13, 15, 20, 20}}, 20))
}
