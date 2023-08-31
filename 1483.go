package partice

import "fmt"

type TreeAncestor struct {
	ancestor map[int][]int // node:depth:ancestor
}

// 超时未通过
func ConstructorTreeAncestor(n int, parent []int) TreeAncestor {
	son := make(map[int]int)
	ancestor := make(map[int][]int)
	for i := 0; i < n; i++ {
		cur := i
		curRes := make([]int, 0)
		dep := 1
		for parent[cur] != -1 {
			if v, ok := son[parent[cur]]; ok {
				ancestor[cur] = ancestor[v]
			}

			curRes = append(curRes, parent[cur])
			dep++
			cur = parent[cur]
		}
		ancestor[i] = curRes
		if _, ok := son[parent[i]]; !ok {
			son[parent[i]] = i
		}
	}
	fmt.Println(ancestor)
	return TreeAncestor{ancestor: ancestor}
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	v, ok := this.ancestor[node]
	if !ok {
		return -1
	}
	if len(v) < k {
		return -1
	}
	return v[k-1]
}
