package partice

func maximalSquare(matrix [][]byte) int {
	m:=len(matrix)
	n:=len(matrix[0])
	ary:=make([][]int,m+1)
	for i:=0;i<=m;i++{
		ary[i]=make([]int,n+1)
	}
	var tmp int
	for i:=1;i<=m;i++{
		for j:=1;j<=n;j++{
			if matrix[i-1][j-1]=='0'{
				continue
			}
			ary[i][j]=min(ary[i-1][j-1],min(ary[i-1][j],ary[i][j-1]))+1
			if ary[i][j]>tmp{
				tmp=ary[i][j]
			}
		}
	}
	return tmp*tmp
}
