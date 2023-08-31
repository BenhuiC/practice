package partice
func runningSum(nums []int) []int {
	s:=0
	for i,v:=range nums{
		s+=v
		nums[i]=s
	}
	return nums
}
