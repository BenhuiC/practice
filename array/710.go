package array

import "math/rand"

type Solution1 struct {
	bound int         // 在n个树中，有bound个非黑名单值
	b2w   map[int]int // 黑名单到白名单的映射，在[0,n)中，把[0,bound)中的黑名单映射到[bound,n)的白名单
}

func Constructor5(n int, blacklist []int) Solution1 {
	m := len(blacklist)
	bound := n - m
	black := make(map[int]bool)
	for _, v := range blacklist {
		if v >= bound {
			black[v] = true
		}
	}
	b2w := make(map[int]int, m-len(black))
	w := bound
	for _, v := range blacklist {
		if v >= bound {
			continue
		}
		for black[w] {
			w++
		}
		b2w[v] = w
		w++
	}
	return Solution1{
		bound: bound,
		b2w:   b2w,
	}
}

func (this *Solution1) Pick() int {
	r := rand.Intn(this.bound)
	if v := this.b2w[r]; v != 0 {
		return v
	}
	return r
}
