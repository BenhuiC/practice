package partice

func countSquares(matrix [][]int) int {
	m:=len(matrix)
	if m==0{
		return 0
	}
	n:=len(matrix[0])
	tp:=make([][]int,m+1)
	for i:=0;i<m+1;i++{
		tp[i]=make([]int,n+1)
	}
	var res1277 int
	for i:=1;i<=m;i++{
		for j:=1;j<=n;j++{
			if matrix[i-1][j-1]!=1{
				continue
			}
			tp[i][j]=min(tp[i-1][j-1],min(tp[i-1][j],tp[i][j-1]))+1
			res1277+=tp[i][j]
		}
	}

	return res1277
}
