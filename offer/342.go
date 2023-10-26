package offer

func isPowerOfFour(n int) bool {
	// 只有一位1并且1在偶数位上
	return n > 0 && n&(n-1) == 0 && n&0xaaaaaaaa == 0
}
