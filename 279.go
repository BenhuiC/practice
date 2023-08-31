package partice

import "math"

func numSquares(n int) int {
	ary:=make([]int,n+1)
	for i:=1;i<=n;i++{
		m:=math.MaxInt32
		for j:=1;j*j<=i;j++{
			m=min(m,ary[i-j*j])
		}
		ary[i]=m+1
	}

	return ary[n]
}

