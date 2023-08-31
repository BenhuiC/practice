package partice

import "fmt"

var tmpMap map[string]float64

func init()  {
	tmpMap=make(map[string]float64)
}

func new21Game(n int, k int, maxPts int) float64 {
	if k==0{
		return 1
	}
	if k==1{
		v:=float64(n-k+1)/float64(maxPts)
		if v>1{
			v=1
		}
		tmpMap[fmt.Sprintf("%d,%d",n,k)]=v
		return v
	}
	if v,ok:=tmpMap[fmt.Sprintf("%d,%d",n,k)];ok{
		return v
	}
	base:=float64(n-k+1)/float64(maxPts)
	for i:=1;i<k;i++{
		base+=0.1*new21Game(n-i,k-i,maxPts)
	}
	return base
}
