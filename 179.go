package partice

import (
	"fmt"
	"sort"
	"strings"
)

func largestNumber(nums []int) string {
	if len(nums)==0{
		return ""
	}
	if len(nums)<=1{
		return fmt.Sprint(nums[0])
	}
	sort.Slice(nums, func(i, j int) bool {
		v1:=fmt.Sprint(nums[i])
		v2:=fmt.Sprint(nums[j])
		return comps(v1,v2)
	})
	res179:=""
	for _,v:=range nums{
		res179=fmt.Sprintf("%s%d",res179,v)
	}
	res179=strings.TrimLeft(res179,"0")
	if len(res179)==0{
		return "0"
	}
	return  res179
}

func comps(v1,v2 string) bool{
	if len(v1)==0{
		return false
	}
	if len(v2)==0{
		return true
	}
	var t int
	for ;t<len(v1)&&t<len(v2);t++{
		if v1[t]<v2[t]{
			return false
		}else if v1[t]>v2[t]{
			return true
		}
	}
	if len(v1)==len(v2){
		return false
	}
	if t==len(v1){
		return comps(v1,v2[t:])
	}
	return comps(v1[t:],v2)
}