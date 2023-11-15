package offer

import "math"

func find132pattern(nums []int) bool {
	n := len(nums)
	candidateK := []int{nums[n-1]}
	maxK := math.MinInt64 // 索引大于当前的最大值

	for i := n - 2; i >= 0; i-- {
		if nums[i] < maxK {
			return true
		}
		// candidateK中的值都大于maxK
		for len(candidateK) > 0 && nums[i] > candidateK[len(candidateK)-1] {
			maxK = candidateK[len(candidateK)-1]
			candidateK = candidateK[:len(candidateK)-1]
		}
		if nums[i] > maxK {
			candidateK = append(candidateK, nums[i])
		}
	}

	return false
}
