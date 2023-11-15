package offer

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	mg := make(map[int]int)
	res := 0
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			mg[v1+v2]++
		}
	}
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			res += mg[-v3-v4]
		}
	}

	return res
}
