package partice

import "sort"

func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	res := make([]int, n)
	n1 := make([]int, n)
	n2 := make([]int, n)
	for i := 0; i < n; i++ {
		n1[i] = i
		n2[i] = i
	}
	sort.Slice(n1, func(i, j int) bool {
		return nums1[n1[i]] < nums1[n1[j]]
	})
	sort.Slice(n2, func(i, j int) bool {
		return nums2[n2[i]] < nums2[n2[j]]
	})

	left, right := 0, n-1
	for i := 0; i < n; i++ {
		if nums1[n1[i]] > nums2[n2[left]] {
			res[n2[left]] = nums1[n1[i]]
			left++
		} else {
			res[n2[right]] = nums1[n1[i]]
			right--
		}
	}

	return res
}
