package partice

func isInterleave(s1 string, s2 string, s3 string) bool {
	m,n,t:=len(s1),len(s2),len(s3)
	if m+n!=t{
		return false
	}
	ary:=make([][]bool,m+1)
	for i:=0;i<m+1;i++{
		ary[i]=make([]bool,n+1)
	}
	ary[0][0]=true
	for i:=0;i<=m;i++{
		for j:=0;j<=n;j++{
			p:=i+j-1
			if i>0{
				ary[i][j]=ary[i][j]||s1[i-1]==s3[p]&&ary[i-1][j]
			}
			if j>0{
				ary[i][j]=ary[i][j]||s2[j-1]==s3[p]&&ary[i][j-1]
			}
		}
	}

	return ary[m][n]
}
