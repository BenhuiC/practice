package partice

func plusOne(digits []int) []int {
	if len(digits)==0{
		return []int{1}
	}
	jin:=1
	for i:=len(digits)-1;i>=0;i--{
		v:=digits[i]+jin
		jin=v/10
		if v<10{
			digits[i]=v
			break
		}
		digits[i]=v%10
	}
	if jin!=0{
		res:=make([]int,0,len(digits)+1)
		res=append([]int{jin},digits...)
		return res
	}
	return digits
}
