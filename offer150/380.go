package offer150

import "math/rand/v2"

type RandomizedSet struct {
	arr []int
	m   map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{arr: []int{}, m: make(map[int]int)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.arr = append(this.arr, val)
	this.m[val] = len(this.arr) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	i, ok := this.m[val]
	if !ok {
		return false
	}
	if i != len(this.arr)-1 {
		this.arr[i] = this.arr[len(this.arr)-1]
		this.m[this.arr[i]] = i
	}
	this.arr = this.arr[:len(this.arr)-1]
	delete(this.m, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	idx := rand.IntN(len(this.arr))
	return this.arr[idx]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
