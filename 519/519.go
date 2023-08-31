package main

import (
	"math/rand"
)

type Solution struct {
	m, n  int
	clean []int
}

func Constructor(m int, n int) Solution {
	s := Solution{
		m: m,
		n: n,
	}
	s.Reset()
	return s
}

func (this *Solution) Flip() []int {
	if len(this.clean) == 0 {
		return nil
	}
	idx := rand.Intn(len(this.clean))
	res := make([]int, 2)
	res[0] = this.clean[idx] / this.n
	res[1] = this.clean[idx] % this.n
	copy(this.clean[idx:], this.clean[idx+1:])
	this.clean = this.clean[:len(this.clean)-1]
	return res
}

func (this *Solution) Reset() {
	if this.clean == nil {
		this.clean = make([]int, this.m*this.n)
		for i := 0; i < this.m*this.n; i++ {
			this.clean[i] = i
		}
		return
	}
	this.clean = this.clean[:0]
	for i := 0; i < this.m*this.n; i++ {
		this.clean = append(this.clean, i)
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(m, n);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
