package array

import "fmt"

func sumSubarrayMins(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	leftMinCnt := make([]int, n)
	rightMinCnt := make([]int, n)
	for i := 0; i < n; i++ {
		l, r := i, i
		for l > 0 && arr[l-1] >= arr[i] {
			l--
		}
		for r < n-1 && arr[r+1] > arr[i] {
			r++
		}
		leftMinCnt[i] = i - l
		rightMinCnt[i] = r - i
	}

	mod := 1000000007
	res := 0
	for i := 0; i < n; i++ {
		mul := leftMinCnt[i] + rightMinCnt[i] + 1 + leftMinCnt[i]*rightMinCnt[i]
		res += mul * arr[i]
		res = res % mod
	}
	return res
}

func sumSubarrayMins2(arr []int) (ans int) {
	const mod int = 1e9 + 7
	n := len(arr)
	left := make([]int, n)
	right := make([]int, n)
	monoStack := []int{}
	for i, x := range arr {
		for len(monoStack) > 0 && x <= arr[monoStack[len(monoStack)-1]] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			left[i] = i + 1
		} else {
			left[i] = i - monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	monoStack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(monoStack) > 0 && arr[i] < arr[monoStack[len(monoStack)-1]] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			right[i] = n - i
		} else {
			right[i] = monoStack[len(monoStack)-1] - i
		}
		monoStack = append(monoStack, i)
	}

	fmt.Println(left, right)
	for i, x := range arr {
		fmt.Printf("res+=%d*%d \t", left[i]*right[i], arr[i])
		ans = (ans + left[i]*right[i]*x) % mod
	}
	fmt.Println()
	return
}
