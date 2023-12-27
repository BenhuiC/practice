package _struct

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestNewSkipList(t *testing.T) {
	s := NewSkipList(4)
	for i := 0; i < 1000; i++ {
		s.Add(rand.Intn(10000))
	}
	fmt.Println(s)

}

func Test_skilList_Search(t *testing.T) {
	s := NewSkipList(4)
	for i := 0; i < 10; i++ {
		s.Add(i)
		fmt.Println(s.Search(i))
	}
	fmt.Println(s)
}

func Test_skilList_Delete(t *testing.T) {
	s := NewSkipList(4)
	for i := 0; i < 100; i++ {
		s.Add(i)
	}
	fmt.Println(s)
	s.Delete(50)
	fmt.Println(s)
}
