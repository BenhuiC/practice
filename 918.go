package partice

import "math"

func maxSubarraySumCircular(nums []int) int {
	if len(nums)==1{
		return nums[0]
	}
	ln:=len(nums)
	tmp:=make([]int,ln)
	res918:=math.MinInt32
	for i:=0;i<ln;i++{
		for j:=i+1;j<ln+i+1;j++{
			ind:=j-i-1
			ind2:=(j+ln)%ln
			if ind==0{
				tmp[ind]=nums[ind2]
				continue
			}
			tmp[ind]=max(nums[ind2],tmp[ind-1]+nums[ind2])
			if tmp[ind]>res918{
				res918=tmp[ind]
			}
		}
	}

	return res918
}
