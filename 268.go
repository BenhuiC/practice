package partice

func missingNumber(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	n := (len(nums) + 1) * len(nums) / 2
	return n - sum
}
