package partice

func twoSum(nums []int, target int) []int {
	res:=make([]int,2)
	i,j:=0,len(nums)-1
	for ;i<j;{
		if nums[i]+nums[j]<target{
			i++
		}else if nums[i]+nums[j]>target{
			j--
		}else {
			res[0]=nums[i]
			res[1]=nums[j]
			return res
		}
	}
	return res
}
