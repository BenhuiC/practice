package _struct

import (
	"math/rand"
	"time"
)

type linkedSkipNode[T any] struct {
	score   int
	val     *T
	forward *linkedSkipNode[T]
	down    *linkedSkipNode[T]
}

type linkedSkipNodeHeader[T any] struct {
	forward *linkedSkipNode[T]
	down    *linkedSkipNodeHeader[T]
	up      *linkedSkipNodeHeader[T]
}

type LinkedSkipList[T any] struct {
	header *linkedSkipNodeHeader[T]
	level  int
}

func NewLinkedSkipList[T any](level int) *LinkedSkipList[T] {
	if level > maxLevel {
		level = maxLevel
	}
	if level <= 0 {
		level = defaultLevel
	}
	list := &LinkedSkipList[T]{level: level}
	var last *linkedSkipNodeHeader[T]
	for i := level; i >= 0; i-- {
		n := &linkedSkipNodeHeader[T]{
			down: last,
		}
		if last != nil {
			last.up = n
		}
		last = n
	}
	list.header = last
	return list
}

func (l *LinkedSkipList[T]) Add(score int, val *T) bool {
	if _, exist := l.Search(score); exist {
		return false
	}
	newLevel := l.randomLevel()
	newNode := linkedSkipNode[T]{
		score:   score,
		val:     val,
		forward: nil,
		down:    nil,
	}
	_ = newNode
	if l.level < newLevel {
		//todo
	}
	for curHeader := l.header; curHeader != nil; curHeader = curHeader.down {

	}
	return true
}

func (l *LinkedSkipList[T]) randomLevel() int {
	// 这里随机数种子需要使用UnixNano，如果使用Unix会导致在同一秒的请求生成的level相同
	rander := rand.New(rand.NewSource(time.Now().UnixNano()))
	lv := 1
	for lv < maxLevel && rander.Int()&1 == 1 {
		lv++
	}
	return lv
}

func (l *LinkedSkipList[T]) Search(score int) (*T, bool) {
	return nil, false
}

func (l *LinkedSkipList[T]) Delete(score int) bool {
	return false
}

func (l *LinkedSkipList[T]) String() string {
	return ""
}
