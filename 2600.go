package partice

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	res2600 := 0
	if k <= numOnes {
		return k
	}
	res2600 = numOnes
	k -= numOnes
	if k <= numZeros {
		return res2600
	}
	k -= numZeros
	res2600 -= k

	return res2600
}
