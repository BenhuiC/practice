package partice

import "strconv"

var m869 = map[int][]string{
	1: {"1", "2", "4", "8"},
	2: {"16", "32", "64"},
	3: {"128", "256", "512"},
	4: {"1024", "2048", "4096", "8192"},
	5: {"16384", "32768", "65536"},
	6: {"131072", "262144", "524288"},
	7: {"1048576", "2097152", "4194304", "8388608"},
	8: {"16777216", "33554432", "67108864"},
	9: {"134217728", "268435456", "536870912"},
}

func reorderedPowerOf2(n int) bool {
	strN := strconv.Itoa(n)
	arys := m869[len(strN)]
	for _, v := range arys {
		if nuEq(strN, v) {
			return true
		}
	}
	return false
}

func nuEq(a, b string) bool {
	if a == b {
		return true
	}
	m1, m2 := map[uint8]int{}, map[uint8]int{}
	for i := range a {
		m1[a[i]]++
		m2[b[i]]++
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}
