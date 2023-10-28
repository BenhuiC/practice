package offer

func intersection(nums1 []int, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	if n1 > n2 {
		nums1, nums2 = nums2, nums1
		n1, n2 = n2, n1
	}
	mp := make(map[int]struct{})
	for _, v := range nums1 {
		mp[v] = struct{}{}
	}
	res := make([]int, 0)
	for _, v := range nums2 {
		if _, ok := mp[v]; ok {
			res = append(res, v)
			delete(mp, v)
		}
	}
	return res
}
