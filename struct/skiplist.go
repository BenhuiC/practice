package _struct

import (
	"fmt"
	"math/rand"
	"time"
)

const maxLevel = 32
const defaultLevel = 4

type skipListNode[T any] struct {
	score   int
	val     *T
	forward []*skipListNode[T]
}

type skipList[T any] struct {
	level  int
	header *skipListNode[T]
}

func NewSkipList[T any](level int) *skipList[T] {
	if level > maxLevel {
		level = maxLevel
	}
	if level <= 0 {
		level = defaultLevel
	}
	return &skipList[T]{
		level:  level,
		header: &skipListNode[T]{score: -1, forward: make([]*skipListNode[T], maxLevel)},
	}
}

func (s *skipList[T]) Add(score int, val *T) bool {
	if _, exist := s.Search(score); exist {
		return false
	}
	newLevel := s.randomLevel()
	if s.level < newLevel {
		s.level = newLevel
	}
	newNode := &skipListNode[T]{
		score:   score,
		val:     val,
		forward: make([]*skipListNode[T], newLevel),
	}
	cur := s.header
	for i := s.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].score < score {
			cur = cur.forward[i]
		}
		if i < newLevel {
			newNode.forward[i] = cur.forward[i]
			cur.forward[i] = newNode
		}
	}
	return true
}

func (s *skipList[T]) randomLevel() int {
	// 这里随机数种子需要使用UnixNano，如果使用Unix会导致在同一秒的请求生成的level相同
	rander := rand.New(rand.NewSource(time.Now().UnixNano()))
	lv := 1
	for lv < maxLevel && rander.Int()&1 == 1 {
		lv++
	}
	return lv
}

func (s *skipList[T]) Search(score int) (val *T, exit bool) {
	cur := s.header
	for i := s.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].score <= score {
			cur = cur.forward[i]
		}
	}
	if cur.score == score {
		return cur.val, true
	}
	return nil, false
}

func (s *skipList[T]) Delete(score int) bool {
	cur := s.header
	isDelete := false
	for i := s.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].score < score {
			cur = cur.forward[i]
		}
		if cur.forward[i] != nil && cur.forward[i].score == score {
			cur.forward[i] = cur.forward[i].forward[i]
			isDelete = true
		}
	}
	return isDelete
}

func (s *skipList[T]) String() string {
	res := fmt.Sprintf("level:%d\n", s.level)
	for i := s.level - 1; i >= 0; i-- {
		cur := s.header
		if cur.forward[i] == nil {
			continue
		}
		res += fmt.Sprintf("header\t")
		for cur = cur.forward[i]; cur != nil; cur = cur.forward[i] {
			res += fmt.Sprintf("\t%d", cur.score)
		}
		res += "\n"
	}

	return res
}
