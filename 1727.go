package partice

import "fmt"


// todo
func largestSubmatrix(matrix [][]int) int {
	if len(matrix)==0||len(matrix[0])==0{
		return 0
	}
	m:=len(matrix)
	n:=len(matrix[0])
	heights:=make([]int,n)
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if matrix[i][j]==0{
				heights[j]=0
			}else{
				heights[j]=heights[j]+1
			}
		}

		fmt.Println(heights)
	}


	return 0
}