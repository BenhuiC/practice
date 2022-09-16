package partice

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
	res1019 := make([]int, 0)
	dec := make([]nodeWithIndex, 0) // 单调递减队列
	idx := 0
	for tmp := head; tmp != nil; tmp = tmp.Next {
		res1019 = append(res1019, 0)
		if len(dec) == 0 {
			dec = append(dec, nodeWithIndex{Idx: idx, Val: tmp.Val})
			idx++
			continue
		}
		if dec[len(dec)-1].Val > tmp.Val {
			dec = append(dec, nodeWithIndex{Idx: idx, Val: tmp.Val})
			idx++
			continue
		}
		for len(dec) > 0 && dec[len(dec)-1].Val < tmp.Val {
			res1019[dec[len(dec)-1].Idx] = tmp.Val
			dec = dec[:len(dec)-1]
		}
		dec = append(dec, nodeWithIndex{Val: tmp.Val, Idx: idx})
		idx++
	}

	return res1019
}

type nodeWithIndex struct {
	Val int //存储节点
	Idx int // 存储对应数组的index
}
