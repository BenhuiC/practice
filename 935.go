package partice

var (
	yu  = 1000000007
	ary = [][]int{{4, 6}, {6, 8}, {7, 9}, {4, 8}, {0, 3, 9}, {}, {0, 1, 7}, {2, 6}, {1, 3}, {2, 4}}
	mp  = map[int]map[int]int{}
)

func init() {
	for i := 0; i < 10; i++ {
		mp[i] = make(map[int]int)
	}
}

func knightDialer(n int) int {
	if n == 1 {
		return 10
	}
	res935 := 0
	for i := range ary {
		res935 += knightF(i, n-2)
		res935 %= yu
	}
	return res935
}

func knightF(idx, n int) int {
	if n == 0 {
		return len(ary[idx])
	}
	if v, ok := mp[idx][n]; ok {
		return v
	}
	res := 0
	for _, v := range ary[idx] {
		res += knightF(v, n-1)
		res = res % yu
	}
	mp[idx][n] = res
	return res
}
