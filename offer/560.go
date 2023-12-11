package offer

func subarraySum(nums []int, k int) int {
	var res, sum int
	mp := make(map[int]int)
	mp[0] = 1
	for _, v := range nums {
		sum += v
		res += mp[sum-k]
		mp[sum]++
	}

	return res
}
