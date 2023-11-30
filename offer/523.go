package offer

func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	m := map[int]int{0: -1}
	var sum int
	for i, v := range nums {
		sum += v
		y := sum % k
		if pidx, has := m[y]; has {
			if i-pidx >= 2 {
				return true
			}
		} else {
			m[y] = i
		}
	}

	return false
}

// timeout
func checkSubarraySum2(nums []int, k int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	leftSum := make([]int, n)
	for i, v := range nums {
		if i == 0 {
			leftSum[i] = v
		} else {
			leftSum[i] = v + leftSum[i-1]
		}
		if i >= 1 && leftSum[i]%k == 0 {
			return true
		}
		for j := 0; j < i-1; j++ {
			diff := leftSum[j] - leftSum[i]
			if diff%k == 0 {
				return true
			}
		}
	}

	return false
}
