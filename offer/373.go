package offer

import "fmt"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n1, n2 := len(nums1), len(nums2)
	m1, m2 := make(map[int]int, n1), make(map[int]int, n2)

	res := make([][]int, 0)
	queue := make([][]int, 0)
	queue = append(queue, []int{0, 0})
	for len(queue) > 0 && len(res) < k {
		idx := 0
		for i := 1; i < len(queue); i++ {
			if nums1[queue[idx][0]]+nums2[queue[idx][1]] > nums1[queue[i][0]]+nums2[queue[i][1]] {
				idx = i
			}
		}
		x, y := queue[idx][0], queue[idx][1]
		m1[x] = y
		m2[y] = x
		res = append(res, []int{nums1[x], nums2[y]})
		queue[idx] = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if y+1 < n2 {
			if x == 0 {
				queue = append(queue, []int{x, y + 1})
			} else if v, ok := m2[y+1]; ok && v == x-1 {
				queue = append(queue, []int{x, y + 1})
			}
		}
		if x+1 < n1 {
			if y == 0 {
				queue = append(queue, []int{x + 1, y})
			} else if v, ok := m1[x+1]; ok && v == y-1 {
				queue = append(queue, []int{x + 1, y})
			}
		}
	}
	return res
}
