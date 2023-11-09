package offer

type lNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *lNode
	TopRight    *lNode
	BottomLeft  *lNode
	BottomRight *lNode
}

func construct(grid [][]int) *lNode {
	var constructPart func(x, y, offset int) *lNode
	constructPart = func(x, y, offset int) *lNode {
		if offset == 1 {
			return &lNode{
				Val:    grid[x][y] == 1,
				IsLeaf: true,
			}
		}
		for i := x; i < x+offset; i++ {
			for j := y; j < y+offset; j++ {
				if grid[i][j] != grid[x][y] {
					return &lNode{
						Val:         false,
						IsLeaf:      false,
						TopLeft:     constructPart(x, y, offset/2),
						TopRight:    constructPart(x, y+offset/2, offset/2),
						BottomLeft:  constructPart(x+offset/2, y, offset/2),
						BottomRight: constructPart(x+offset/2, y+offset/2, offset/2),
					}
				}
			}
		}
		return &lNode{
			Val:    grid[x][y] == 1,
			IsLeaf: true,
		}
	}
	return constructPart(0, 0, len(grid))
}
