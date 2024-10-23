package array

func minOperations(nums []int) int {
	res := 0
	n := len(nums)
	for i, v := range nums {
		if v == 1 {
			continue
		}
		if i > n-3 {
			return -1
		}
		nums[i] ^= 1
		nums[i+1] ^= 1
		nums[i+2] ^= 1
		res++
	}

	return res
}
