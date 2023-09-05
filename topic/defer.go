package main

import "fmt"

/*
defer，return的执行顺序
先执行return，对于无名的返回值，return会将返回的值写入返回的变量中。再执行defer。
*/

func deferWithoutName() int {
	i := 1
	defer func() {
		i++
		fmt.Println("deferWithoutName i is ", i)
	}()
	return i
}

func deferWithName() (i int) {
	i = 1
	defer func() {
		i++
		fmt.Println("deferWithName i is ", i)
	}()
	return i
}

func deferWithPoint() *int {
	i := 1
	defer func() {
		i++
		fmt.Println("deferWithPoint i is ", i)
	}()
	return &i
}
