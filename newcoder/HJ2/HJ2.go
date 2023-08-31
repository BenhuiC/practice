package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _, _ := reader.ReadLine()
	tgt, _, _ := reader.ReadLine()
	m := make(map[byte]int)
	for _, v := range str {
		m[v]++
	}
	res := 0
	if tgt[0] >= 'A' && tgt[0] <= 'Z' {
		res = m[tgt[0]] + m[tgt[0]+32]
	} else if tgt[0] >= 'a' && tgt[0] <= 'z' {
		res = m[tgt[0]] + m[tgt[0]-32]
	} else {
		res = m[tgt[0]]
	}
	fmt.Println(res)
}
