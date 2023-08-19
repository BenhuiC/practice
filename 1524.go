package partice

func numOfSubarrays(arr []int) int {
	if len(arr) <= 0 {
		return 0
	}
	mod := 1000000007
	res1524 := 0
	odd, even := 0, 0 // 使用前缀和的思路，记录到当前数为止，和为奇数偶数的数量
	cur := 0
	for _, v := range arr {
		cur = v&1 ^ cur
		if cur == 1 {
			// 如果当前的和为奇数，则以当前数字为尾的奇数和区间为偶数数量+1（奇数-偶数=奇数，再加上当前数字）
			res1524 = (res1524 + even + 1) % mod
			odd++
		} else {
			// 如果为偶数，则以当前数字为为的奇数区间为奇数数量
			res1524 = (res1524 + odd) % mod
			even++
		}
	}
	return res1524
}
