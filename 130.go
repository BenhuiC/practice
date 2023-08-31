package partice

import "fmt"

func solve(board [][]byte)  {
	m,n:=len(board),len(board[0])
	for i:=0;i<len(board);i++{
		if board[i][0]=='O'{
			oo(board,i,0)
		}
		if board[i][n-1]=='O'{
			oo(board,i,n-1)
		}
	}
	for j:=0;j<n;j++{
		if board[0][j]=='O'{
			oo(board,0,j)
		}
		if board[m-1][j]=='O'{
			oo(board,m-1,j)
		}
	}
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if board[i][j]=='o'{
				board[i][j]='O'
			}else{
				board[i][j]='X'
			}
		}
	}
}

func oo(b [][]byte,x,y int){
	type tp struct {
		x int
		y int
	}
	lr:=[]tp{{-1,0},{1,0},{0,-1},{0,1}}
	st:=make([]tp,0)
	st = append(st, tp{x,y})
	m:=make(map[string]bool)
	m[fmt.Sprintf("%d,%d",x,y)]=true
	for {
		if len(st)==0{
			break
		}
		x,y=st[0].x,st[0].y
		b[x][y]='o'
		for _,v:=range lr{
			tx:=x+v.x
			ty:=y+v.y
			if tx>=0&&tx<len(b)&&ty>=0&&ty<len(b[0])&&!m[fmt.Sprintf("%d,%d",tx,ty)]&&b[tx][ty]=='O'{
				st = append(st, tp{tx,ty})
				m[fmt.Sprintf("%d,%d",tx,ty)]=true
			}
		}
		st=st[1:]
	}
}