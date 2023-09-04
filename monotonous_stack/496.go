package monotonous_stack

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nextGreater := make(map[int]int, len(nums2))
	ary := make([]int, 0)
	ary = append(ary, nums2[0])
	idx := 1
	for len(ary) != 0 && idx < len(nums2) {
		if ary[len(ary)-1] > nums2[idx] {
			ary = append(ary, nums2[idx])
			idx++
			continue
		}
		for len(ary) != 0 && ary[len(ary)-1] < nums2[idx] {
			nextGreater[ary[len(ary)-1]] = nums2[idx]
			ary = ary[:len(ary)-1]
		}
		ary = append(ary, nums2[idx])
		idx++
	}

	res := make([]int, len(nums1))
	for i, v := range nums1 {
		if g, ok := nextGreater[v]; ok {
			res[i] = g
		} else {
			res[i] = -1
		}
	}
	return res
}
