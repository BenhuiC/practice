package hot

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	l, r := -1, -1
	aidx, bidx := 0, 0
	for i := 0; i <= (m+n)/2; i++ {
		l = r
		if aidx < m && (bidx < n && nums1[aidx] < nums2[bidx] || bidx >= n) {
			r = nums1[aidx]
			aidx++
		} else {
			r = nums2[bidx]
			bidx++
		}
	}

	if (m+n)&1 == 1 {
		return float64(r)
	} else {
		return float64(l+r) / 2
	}
}
