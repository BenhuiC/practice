package partice

func mySqrt(x int) int {
	res := 1
	for ; (res+1)*(res+1) < x; res++ {
	}
	return res
}
