package hot

func subarraySum(nums []int, k int) int {
	var res, sum int
	m := make(map[int]int)
	m[0] = 1
	for _, v := range nums {
		sum += v
		res += m[sum-k]
		m[sum]++
	}
	return res
}
