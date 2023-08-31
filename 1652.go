package partice

func decrypt(code []int, k int) []int {
	res := make([]int, len(code))
	n := len(code)
	if k == 0 {
		return res
	}
	ksum := 0
	if k > 0 {
		for i := 0; i < k; i++ {
			ksum += code[i]
		}
		for i, v := range code {
			ksum -= v
			ksum += code[(i+k)%n]
			res[i] = ksum
		}
		return res
	}
	k = -k
	for i := 0; i < k; i++ {
		ksum += code[n-i-1]
	}
	for i := n - 1; i >= 0; i-- {
		ksum -= code[i]
		ksum += code[(i-k+n)%n]
		res[i] = ksum
	}

	return res
}
