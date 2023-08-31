package partice

func canThreePartsEqualSum(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	allSum := 0
	for _, v := range arr {
		allSum += v
	}
	if allSum%3 != 0 {
		return false
	}
	avg := allSum / 3
	leftSum, rightSum := 0, 0
	for i := 0; i < len(arr)-2; i++ {
		leftSum += arr[i]
		rightSum = 0
		if leftSum == avg {
			for j := len(arr) - 1; j-i > 1; j-- {
				rightSum += arr[j]
				if rightSum == avg {
					return true
				}
			}
		}
	}

	return false
}
