package partice

func balancedStringSplit(s string) int {
	var res1221 int
	var l,r int
	for i:=0;i<len(s);i++{
		if s[i]=='L'{
			l++
		}else {
			r++
		}
		if l==r{
			l=0
			r=0
			res1221++
		}
	}

	return res1221
}