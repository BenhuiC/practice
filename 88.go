package partice

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	copy(nums1[len(nums1)-m:], nums1[:m])
	idx, i1, i2 := 0, n, 0
	for idx = 0; idx < len(nums1); idx++ {
		if i1 >= len(nums1) {
			nums1[idx] = nums2[i2]
			i2++
			continue
		}
		if i2 >= n {
			nums1[idx] = nums1[i1]
			i1++
			continue
		}
		if nums1[i1] < nums2[i2] {
			nums1[idx] = nums1[i1]
			i1++
		} else {
			nums1[idx] = nums2[i2]
			i2++
		}
	}
}
