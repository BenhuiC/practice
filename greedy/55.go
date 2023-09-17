package greedy

func canJump(nums []int) bool {
	can := make([]bool, len(nums))
	can[0] = true
	for i := 0; i < len(nums); i++ {
		if !can[i] {
			break
		}
		for j := i; j < len(nums) && j <= nums[i]+i; j++ {
			can[j] = true
		}
	}

	return can[len(nums)-1]
}

func canJump2(nums []int) bool {
	/*
		n表示为到达上一个标记点的步数。对于点j，如果当前点i，i到j的距离为n，当前节点i的步长nums[i]
		如果nums[i]>=n，则从i到j只需要一步
	*/
	n := 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= n {
			n = 1
		} else {
			n++
		}
	}

	return n == 1
}
