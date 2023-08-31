package interview

var dp1709 []int

func init() {
	dp1709 = make([]int, 1001)
	dp1709[1] = 1
	p3, p5, p7 := 1, 1, 1
	for i := 2; i <= 1000; i++ {
		x3, x5, x7 := dp1709[p3]*3, dp1709[p5]*5, dp1709[p7]*7
		dp1709[i] = Min(x3, Min(x5, x7))
		if dp1709[i] == x3 {
			p3++
		}
		if dp1709[i] == x5 {
			p5++
		}
		if dp1709[i] == x7 {
			p7++
		}
	}
}

func getKthMagicNumber(k int) int {
	return dp1709[k]
}

/*
1
3
5
7

9
15
21
35

27
45
63
105

75
175

147
245




*/
