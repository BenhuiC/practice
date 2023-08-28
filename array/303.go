package array

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	ary := make([]int, len(nums))
	ary[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		ary[i] = ary[i-1] + nums[i]
	}
	return NumArray{prefixSum: ary}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left < 0 {
		return 0
	}
	if right >= len(this.prefixSum) {
		return 0
	}
	l, r := 0, this.prefixSum[right]
	if left != 0 {
		l = this.prefixSum[left-1]
	}
	return r - l
}
