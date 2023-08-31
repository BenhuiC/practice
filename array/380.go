package array

import "math/rand"

type RandomizedSet struct {
	dataMap map[int]int
	data    []int
}

func Constructor4() RandomizedSet {
	return RandomizedSet{
		dataMap: make(map[int]int),
		data:    make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.dataMap[val]; ok {
		return false
	}
	this.data = append(this.data, val)
	this.dataMap[val] = len(this.data) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.dataMap[val]
	if !ok {
		return false
	}
	n := len(this.data)
	this.data[idx] = this.data[n-1]
	this.dataMap[this.data[idx]] = idx
	this.data = this.data[:n-1]
	delete(this.dataMap, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	n := len(this.dataMap)
	return this.data[rand.Intn(n)]
}
