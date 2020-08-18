package partice

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return mid(nums2, len(nums2))
	}
	if len(nums2) == 0 {
		return mid(nums1, len(nums1))
	}
	l1 := len(nums1) + len(nums2)
	m := make([]int, l1)
	n := 0
	i, j := 0, 0
	for {
		for i < len(nums1) && nums1[i] < nums2[j] {
			m[n] = nums1[i]
			i++
			n++
		}
		for j < len(nums2) && nums2[j] <= nums1[i] {
			m[n] = nums2[j]
			j++
			n++
		}
		if n >= l1/2+1 {
			break
		}
	}
	return mid(m, l1)
}

func mid(n []int, length int) float64 {
	if length == 0 {
		return 0
	}

	if length%2 == 1 {
		return float64(n[length/2])
	}
	return (float64(n[length/2-1]) + float64(n[length/2])) / 2.0
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return mid(nums2, len(nums2))
	}
	if len(nums2) == 0 {
		return mid(nums1, len(nums1))
	}
	l1 := len(nums1) + len(nums2)
	m := make([]int, 0)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			m = append(m, nums1[i])
			i++
		} else {
			m = append(m, nums2[j])
			j++
		}
		if len(m) >= l1/2+1 {
			break
		}
	}
	if i == len(nums1) {
		m = append(m, nums2[j:]...)
	} else if j == len(nums2) {
		m = append(m, nums1[i:]...)
	}
	return mid(m, l1)
}
