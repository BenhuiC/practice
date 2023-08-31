package partice

var boardA [][]byte
var ck [][]bool

func exist(board [][]byte, word string) bool {
	var result bool
	if len(board)==0||len(board[0])==0{
		return false
	}
	if len(word)>(len(board)*len(board[0])){
		return false
	}
	boardA=board
	ck=make([][]bool,len(board))
	for i:=range ck{
		ck[i]=make([]bool,len(board[0]))
	}
	for i:=0;i<len(board);i++{
		for j:=0;j<len(board[0]);j++{
			result=meet(i,j,0,word)
			if result{
				return result
			}
		}
	}
	return false
}

func meet(x,y,start int,word string) bool{
	if start>=len(word){
		return true
	}
	if x<0||x>=len(boardA)||y<0||y>=len(boardA[0]){
		return false
	}
	if ck[x][y]{
		return false
	}
	if boardA[x][y]!=word[start]{
		return false
	}
	ck[x][y]=true
	s:=[][]int{{-1,0},{0,1},{1,0},{0,-1}}
	for _,v:=range s{
		if meet(x+v[0],y+v[1],start+1,word){
			ck[x][y]=false
			return true
		}
	}

	ck[x][y]=false
	return false
}

