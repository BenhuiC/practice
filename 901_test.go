package partice

import (
	"fmt"
	"testing"
)

func TestStockSpanner_Next(t *testing.T) {
	S := Constructor901()
	fmt.Println(S.Next(100))
	fmt.Println(S.Next(80))
	fmt.Println(S.Next(60))
	fmt.Println(S.Next(70))
	fmt.Println(S.Next(60))
	fmt.Println(S.Next(75))
	fmt.Println(S.Next(85))
	fmt.Println(S.Next(85))
	fmt.Println(S.Next(85))
	fmt.Println(S.Next(86))
}
