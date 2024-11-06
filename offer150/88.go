package offer150

func merge(nums1 []int, m int, nums2 []int, n int) {
	idx1, idx2 := m-1, n-1
	idx := m + n - 1
	for idx1 >= 0 || idx2 >= 0 {
		if idx1 < 0 {
			nums1[idx] = nums2[idx2]
			idx2--
		} else if idx2 < 0 {
			nums1[idx] = nums1[idx1]
			idx1--
		} else {
			if nums1[idx1] > nums2[idx2] {
				nums1[idx] = nums1[idx1]
				idx1--
			} else {
				nums1[idx] = nums2[idx2]
				idx2--
			}
		}

		idx--
	}
}
