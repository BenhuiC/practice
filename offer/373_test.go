package offer

import (
	"fmt"
	"testing"
)

func Test_kSmallestPairs(t *testing.T) {
	res := kSmallestPairs([]int{1, 3, 7, 11, 11, 23, 24, 30, 100, 101, 202}, []int{2, 4, 6, 20, 40, 100}, 10)
	fmt.Println(res)
}
