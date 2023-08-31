package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var m map[int]bool
var res int
var ji, ou []int
var jiMap, ouMap []bool

func main() {
	reader := bufio.NewReader(os.Stdin)
	m = make(map[int]bool)
	for {
		nu, _, err := reader.ReadLine()
		if err != nil || len(nu) == 0 {
			break
		}
		da, _, _ := reader.ReadLine()
		arys := strings.Split(string(da), " ")
		group(arys)
		res = 0
		backt(0)
		fmt.Println(res)
	}
}

func group(ary []string) {
	ji = make([]int, 0)
	ou = make([]int, 0)
	for _, v := range ary {
		val, _ := strconv.Atoi(v)
		if val%2 == 0 {
			ou = append(ou, val)
		} else {
			ji = append(ji, val)
		}
	}
	jiMap = make([]bool, len(ji))
	ouMap = make([]bool, len(ou))
}

func backt(n int) {
	for i := 0; i < len(ji); i++ {
		if jiMap[i] {
			continue
		}
		jiMap[i] = true
		for j := 0; j < len(ouMap); j++ {
			if ouMap[j] {
				continue
			}
			val := ji[i] + ou[j]
			if m[val] || su(val) {
				m[val] = true
				ouMap[j] = true
				res = max(res, n+1)
				backt(n + 1)
				ouMap[j] = false
			}
		}
		jiMap[i] = false
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func su(n int) bool {
	x := int(math.Sqrt(float64(n)))
	for i := 2; i <= x; i++ {
		if n%x == 0 {
			return false
		}
	}
	return true
}
