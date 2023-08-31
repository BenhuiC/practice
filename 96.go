package partice


func numTrees(n int) int {
	if n<=1{
		return 1
	}
	if n==2{
		return 2
	}

	var tmp int
	for i:=1;i<=n;i++{
		tmp+=numTrees(i-1)*numTrees(n-i)
	}
	return tmp
}
