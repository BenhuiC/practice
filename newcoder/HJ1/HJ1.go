package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	da, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	n := len(da)
	i := n - 1
	for ; i >= 0; i-- {
		if da[i] == ' ' {
			break
		}
	}
	fmt.Println(n - i - 1)
}
