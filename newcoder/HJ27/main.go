package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
3 abc bca cab abc 1
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	da, _, _ := reader.ReadLine()
	strs, tgt, n := parseInput(da)
	broAry := make([]string, 0)
	m := make(map[byte]int)
	for i := range tgt {
		m[tgt[i]]++
	}
	for _, v := range strs {
		if isBro(v, tgt, m) {
			broAry = append(broAry, v)
		}
	}
	sort.Strings(broAry)
	fmt.Println(len(broAry))
	if len(broAry) >= n {
		fmt.Println(broAry[n-1])
	}
}

func parseInput(da []byte) ([]string, string, int) {
	ary := strings.Split(string(da), " ")
	if len(ary) == 0 {
		return nil, "", 0
	}
	l, _ := strconv.Atoi(ary[0])
	strs := ary[1 : 1+l]
	tgt := ary[l+1]
	n, _ := strconv.Atoi(ary[l+2])
	return strs, tgt, n
}

func isBro(a, b string, m map[byte]int) bool {
	if a == b {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	m1 := make(map[byte]int)
	for i := range a {
		m1[a[i]]++
	}
	for k := range m {
		if m1[k] != m[k] {
			return false
		}
	}
	return true
}
