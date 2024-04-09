package array

import "sort"

func threeSumMulti(arr []int, target int) int {
	n := len(arr)
	res := 0
	mod := 1000000007
	sort.Ints(arr)
	for i := 0; i < n-2; i++ {
		t := target - arr[i]
		l, r := i+1, n-1
		for l < r {
			if arr[l]+arr[r] > t {
				r--
			} else if arr[l]+arr[r] < t {
				l++
			} else {
				if arr[l] != arr[r] {
					lNu := 1
					rNu := 1
					for l+1 < n && arr[l+1] == arr[l] {
						lNu++
						l++
					}
					for r-1 >= 0 && arr[r-1] == arr[r] {
						rNu++
						r--
					}
					res += lNu * rNu
					res = res % mod
					l++
					r--
				} else {
					res += (r - l + 1) * (r - l) / 2
					res = res % mod
					break
				}
			}
		}
	}

	return res
}
