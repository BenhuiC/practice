package partice

/*
5  00101
25 11001
28 11100
*/

// todo 超出时间限制
func findMaximumXOR(nums []int) int {
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			xor := nums[i] ^ nums[j]
			if xor > res {
				res = xor
			}
		}
	}
	return res
}
