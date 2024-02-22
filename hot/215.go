package hot

func findKthLargest(nums []int, k int) int {
	bucket := make([]int, 20001)
	for _, v := range nums {
		bucket[v+10000]++
	}
	for i := 20000; i >= 0; i-- {
		k -= bucket[i]
		if k <= 0 {
			return i - 10000
		}
	}
	return 0
}
