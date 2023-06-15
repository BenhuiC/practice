package partice

func canMakePaliQueries(s string, queries [][]int) []bool {
	res1177 := make([]bool, len(queries))
	// 用前缀和的思想，prefix[i][j]表示s[0:i]中字符j出现的次数
	prefix := make([][26]int, len(s)+1)
	prefix[0] = [26]int{}
	for i := 0; i < len(s); i++ {
		prefix[i+1] = [26]int{}
		copy(prefix[i+1][:], prefix[i][:])
		prefix[i+1][s[i]-'a']++
	}

	// 对每个查询，统计s[left:right+1]中每个字符出现的次数
	// 因为可以重新排序，所以偶数次数的字符可以忽略，只需要统计奇数次数的字符
	// 因为可以替换任意字符，所以一对奇数次数的字符可以通过一次替换变为两个偶数次数的字符
	// 所以如果oddCount/2 <= k，则可以通过k次替换将s[left:right+1]变为回文串
	for j, v := range queries {
		left, right, k := v[0], v[1], v[2]
		if left == right {
			res1177[j] = true
			continue
		}
		var oddCount int
		for c := 0; c < 26; c++ {
			diff := prefix[right+1][c] - prefix[left][c]
			if diff&1 == 1 {
				oddCount++
			}
		}
		if oddCount/2 <= k {
			res1177[j] = true
		}
	}

	return res1177
}

func canMakePaliQueries2(s string, queries [][]int) []bool {
	res1177 := make([]bool, len(queries))
	prefix := make([]int, len(s)+1)
	prefix[0] = 0
	for i := 0; i < len(s); i++ {
		prefix[i+1] = prefix[i] ^ (1 << (s[i] - 'a'))
	}

	for j, v := range queries {
		left, right, k := v[0], v[1], v[2]
		if left == right {
			res1177[j] = true
			continue
		}
		var oddCount int
		diff := prefix[right+1] ^ prefix[left]
		for diff > 0 {
			if diff&1 == 1 {
				oddCount++
			}
			diff = diff >> 1
		}
		if oddCount/2 <= k {
			res1177[j] = true
		}
	}

	return res1177
}
