package partice

func removeElement(nums []int, val int) int {
	var i, j int
	for ; j < len(nums); j++ {
		if nums[j] == val {
			continue
		}
		nums[i] = nums[j]
		i++
	}
	return i
}
