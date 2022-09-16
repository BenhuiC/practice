package partice

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for i, v := range nums2 {
		m[v] = -1
		for j := i + 1; j < len(nums2); j++ {
			if nums2[j] > v {
				m[v] = nums2[j]
				break
			}
		}
	}
	for i := range nums1 {
		nums1[i] = m[nums1[i]]
	}
	return nums1
}
