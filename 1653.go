package partice

func minimumDeletions(s string) int {
	n := len(s)
	a := make([]int, n+1)
	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = a[i-1]
		b[i] = b[i-1]
		if s[i-1] == 'a' {
			a[i]++
		} else {
			b[i]++
		}
	}
	res1653 := n
	for i := 1; i <= n; i++ {
		res1653 = min(res1653, b[i-1]+a[n]-a[i])
	}
	return res1653
}
