package array

func subarraysDivByK(nums []int, k int) int {
	sum := 0
	prefix := make(map[int]int)
	res := 0
	prefix[0] = 1
	for _, v := range nums {
		sum += v
		mod := (sum%k + k) % k
		res += prefix[mod]
		prefix[mod]++
	}

	return res
}
