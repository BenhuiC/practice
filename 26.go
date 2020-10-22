package partice

func removeDuplicates(nums []int) int {
	var i, j int
	for {
		for j = i + 1; j < len(nums) && nums[j] == nums[i]; j++ {
		}
		if j-i > 1 {
			t := j - i
			for k := i + 1; k < len(nums)-t+1; k++ {
				nums[k] = nums[k+t-1]
			}
			nums = nums[:len(nums)-t+1]
		}
		i++
		if i >= len(nums) {
			break
		}
	}
	return len(nums)
}

func removeDuplicates2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	var j int
	for i := 1; i < len(nums); i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
