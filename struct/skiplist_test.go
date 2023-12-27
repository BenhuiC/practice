package _struct

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestNewSkipList(t *testing.T) {
	s := NewSkipList[int](4)
	for i := 0; i < 1000; i++ {
		v := i
		s.Add(rand.Intn(10000), &v)
	}
	fmt.Println(s)

}

func Test_skilList_Search(t *testing.T) {
	s := NewSkipList[int](4)
	for i := 0; i < 100; i++ {
		v := i
		s.Add(i, &v)
	}
	fmt.Println(s)
	for i := 0; i < 10; i++ {
		v := rand.Intn(200)
		search, ok := s.Search(v)
		if !ok {
			fmt.Printf("%d not found\n", v)
		} else {
			fmt.Printf("%d values is %v\n", v, *search)
		}
	}
}

func Test_skilList_Delete(t *testing.T) {
	s := NewSkipList[int](4)
	for i := 0; i < 100; i++ {
		v := i
		s.Add(i, &v)
	}
	fmt.Println(s)
	for i := 0; i < 10; i++ {
		v := rand.Intn(200)
		ok := s.Delete(v)
		if ok {
			fmt.Printf("delete %d success\n", v)
		} else {
			fmt.Printf("%d not found\n", v)
		}
	}
	fmt.Println(s)
}
