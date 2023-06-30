package partice

import "math"

func pivotInteger(n int) int {
	target := math.Sqrt(float64(n * (n + 1) / 2))
	if target == float64(int(target)) {
		return int(target)
	}

	return -1
}

/*
x*(1+x)= (n+x)*(n-x+1)
x+x*x=n*n-x*x+n+x
2(x*x)=n*n+n
*/
