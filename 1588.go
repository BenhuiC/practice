package partice



func sumOddLengthSubarrays(arr []int) int {
	var res1588 int
	n:=len(arr)
	for i := 0; i < n; i++ {
		leftCount ,rightCount:= i,   n - i - 1
		leftOdd := (leftCount + 1) / 2
		rightOdd := (rightCount + 1) / 2
		leftEven := leftCount / 2 + 1
		rightEven := rightCount / 2 + 1
		res1588 += arr[i] * (leftOdd * rightOdd + leftEven * rightEven)
	}
	return res1588
}
