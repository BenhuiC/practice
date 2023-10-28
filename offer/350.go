package offer

func intersect(nums1 []int, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	if n1 > n2 {
		nums1, nums2 = nums2, nums1
		n1, n2 = n2, n1
	}
	mp := make(map[int]int)
	for _, v := range nums1 {
		mp[v]++
	}
	res := make([]int, 0)
	for _, v := range nums2 {
		if nu, ok := mp[v]; ok && nu > 0 {
			res = append(res, v)
			mp[v]--
		}
	}
	return res
}
