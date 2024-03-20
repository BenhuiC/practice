package array

import "bytes"

func decodeAtIndex(s string, k int) string {
	type item struct {
		chars []byte
		mul   int
	}
	ary := make([]item, 0)
	total := 0

	chars := make([]byte, 0)
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			chars = append(chars, byte(c))
			continue
		}

		nu := int(c - '0')
		ary = append(ary, item{chars: chars, mul: nu})
		total = (total + len(chars)) * nu
		chars = make([]byte, 0)
		if total < k {
			continue
		}
		for i := len(ary) - 1; i >= 0; i-- {
			total /= ary[i].mul
			for k > total {
				k -= total
			}
			if k > total-len(ary[i].chars) && len(ary[i].chars) > 0 {
				return string(ary[i].chars[k-total+len(ary[i].chars)-1])
			} else {
				total -= len(ary[i].chars)
			}
		}
	}

	if len(chars) >= k {
		return string(chars[k-1])
	}
	return ""
}

// oom
func decodeAtIndex2(s string, k int) string {
	chars := make([]byte, 0)
	for _, c := range s {
		clen := len(chars)
		if clen >= k {
			return string(chars[k-1])
		}
		if c >= 'a' && c <= 'z' {
			chars = append(chars, byte(c))
			continue
		}
		nu := int(c - '0')
		chars = bytes.Repeat(chars, nu)
	}
	if len(chars) >= k {
		return string(chars[k-1])
	}
	return ""
}
