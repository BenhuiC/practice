package partice

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	var res128 int
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	for k := range m {
		if m[k-1] {
			continue
		}
		for t := k; m[t]; t++ {
			if t-k+1 > res128 {
				res128 = t - k + 1
			}
		}
	}
	return res128
}
