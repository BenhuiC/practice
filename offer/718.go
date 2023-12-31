package offer

func findLength(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if nums1[i] != nums2[j] {
				continue
			}
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}
