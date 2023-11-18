package offer

func findMaxConsecutiveOnes(nums []int) int {
	res := 0
	i, j := 0, 0
	for i < len(nums) {
		if nums[i] == 0 {
			i++
			continue
		}
		for j = i + 1; j < len(nums) && nums[j] == 1; j++ {
		}
		tmp := j - i
		if tmp > res {
			res = tmp
		}
		i = j
	}
	return res
}
