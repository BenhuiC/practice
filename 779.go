package partice

func kthGrammar(n int, k int) int {
	//history := []string{"0"}
	//
	//var get func(n, k int) int
	//get = func(n, k int) int {
	//	if len(history[len(history)-1]) >= k {
	//		return int(history[len(history)-1][k-1] - '0')
	//	}
	//	last := history[len(history)-1]
	//	for len(last) < k {
	//		newVal := ""
	//		for t := range last {
	//			if last[t] == '0' {
	//				newVal += "01"
	//			} else {
	//				newVal += "10"
	//			}
	//		}
	//		history = append(history, newVal)
	//		last = newVal
	//	}
	//	return int(last[k-1] - '0')
	//}
	//
	//return get(n, k)

	if k == 1 {
		return 0
	}
	if k > 1<<(n-2) {
		return 1 ^ kthGrammar(n-1, k-1<<(n-2))
	}
	return kthGrammar(n-1, k)

}

/*
0
01
0110
0110 1001
0110100 110010110
0110100110010110 101001011001101001

*/
