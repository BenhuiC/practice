package partice

func rotate189(nums []int, k int) {
	n := len(nums)
	k = k % n
	rev(nums)
	rev(nums[:k])
	rev(nums[k:])
}

func rev(num []int) {
	for i, j := 0, len(num)-1; i < j; {
		num[i], num[j] = num[j], num[i]
		i++
		j--
	}
}
