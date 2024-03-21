package array

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	var aliceSum, bobSum int
	for _, a := range aliceSizes {
		aliceSum += a
	}
	bobMap := make(map[int]bool)
	for _, b := range bobSizes {
		bobSum += b
		bobMap[b] = true
	}

	diff := (bobSum+aliceSum)/2 - aliceSum
	for _, a := range aliceSizes {
		if bobMap[a+diff] {
			return []int{a, a + diff}
		}
	}
	return nil
}
