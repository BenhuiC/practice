package _struct

import (
	"fmt"
	"math/rand"
	"time"
)

const maxLevel = 32
const defaultLevel = 4

type skipListNode struct {
	val     int
	forward []*skipListNode
}

type skilList struct {
	level  int
	header *skipListNode
}

func NewSkipList(level int) *skilList {
	if level > maxLevel {
		level = maxLevel
	}
	if level <= 0 {
		level = defaultLevel
	}
	return &skilList{
		level:  level,
		header: &skipListNode{val: -1, forward: make([]*skipListNode, maxLevel)},
	}
}

func (s *skilList) Add(val int) {
	if s.Search(val) {
		return
	}
	newLevel := s.randomLevel()
	if s.level < newLevel {
		s.level = newLevel
	}
	newNode := &skipListNode{
		val:     val,
		forward: make([]*skipListNode, newLevel),
	}
	cur := s.header
	for i := s.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].val < val {
			cur = cur.forward[i]
		}
		if i < newLevel {
			newNode.forward[i] = cur.forward[i]
			cur.forward[i] = newNode
		}
	}
}

func (s *skilList) randomLevel() int {
	// 这里随机数种子需要使用UnixNano，如果使用Unix会导致在同一秒的请求生成的level相同
	rander := rand.New(rand.NewSource(time.Now().UnixNano()))
	lv := 1
	for lv < maxLevel && rander.Int()&1 == 1 {
		lv++
	}
	return lv
}

func (s *skilList) Search(val int) bool {
	cur := s.header
	for i := s.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].val <= val {
			cur = cur.forward[i]
		}
	}
	if cur.val == val {
		return true
	}
	return false
}

func (s *skilList) Delete(val int) {
	cur := s.header
	for i := s.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].val < val {
			cur = cur.forward[i]
		}
		if cur.forward[i] != nil && cur.forward[i].val == val {
			cur.forward[i] = cur.forward[i].forward[i]
		}
	}
}

func (s *skilList) String() string {
	res := fmt.Sprintf("level:%d\n", s.level)
	for i := s.level - 1; i >= 0; i-- {
		cur := s.header
		if cur.forward[i] == nil {
			continue
		}
		res += fmt.Sprintf("header\t")
		for cur = cur.forward[i]; cur != nil; cur = cur.forward[i] {
			res += fmt.Sprintf("\t%d", cur.val)
		}
		res += "\n"
	}

	return res
}
