package partice

import (
	"testing"
)

func genList(nums []int) (h *ListNode) {
	if len(nums) == 0 {
		return
	}
	h = &ListNode{Val: nums[0]}
	t := h
	for i := 1; i < len(nums); i++ {
		t.Next = &ListNode{Val: nums[i]}
		t = t.Next
	}

	return h
}

func Test_nextLargerNodes(t *testing.T) {
	tests := []struct {
		head *ListNode
	}{
		{genList([]int{3, 2, 1})},
		{genList([]int{2, 1, 5})},
		{genList([]int{2, 7, 4, 3, 5})},
		{genList([]int{1, 7, 5, 1, 9, 2, 5, 1})},
	}
	for _, tt := range tests {
		got := nextLargerNodes(tt.head)
		t.Log(got)
	}
}
