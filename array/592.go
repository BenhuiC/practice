package array

import "fmt"

func fractionAddition(expression string) string {
	type nu struct {
		sign        byte
		molecular   int
		denominator int
	}
	add := func(a, b nu) nu {
		if a.denominator == 0 || a.molecular == 0 {
			return b
		}
		if b.denominator == 0 || b.molecular == 0 {
			return a
		}
		r := nu{}
		r.denominator = a.denominator * b.denominator
		molecularA, molecularB := a.molecular*b.denominator, b.molecular*a.denominator

		if a.sign == b.sign {
			r.sign = a.sign
			r.molecular = molecularA + molecularB
		} else {
			if molecularA > molecularB {
				r.sign = a.sign
				r.molecular = molecularA - molecularB
			} else if molecularA < molecularB {
				r.sign = b.sign
				r.molecular = molecularB - molecularA
			} else {
				r.sign = '+'
				r.denominator = 1
				r.molecular = 0
				return r
			}
		}

		x, y, mod := r.denominator, r.molecular, 0
		for {
			if x < y {
				x, y = y, x
			}
			m := x % y
			if m == 0 {
				mod = y
				break
			}
			x = m
		}
		r.denominator /= mod
		r.molecular /= mod

		return r
	}

	last := nu{sign: '+'}
	cur := nu{sign: '+'}
	v := 0
	for i, c := range expression {
		if c == '+' || c == '-' {
			if i == 0 {
				cur = nu{sign: byte(c)}
			} else {
				cur.denominator = v
				v = 0
				last = add(last, cur)
				cur = nu{sign: byte(c)}
			}
		} else if c == '/' {
			cur.molecular = v
			v = 0
		} else {
			v = v*10 + int(c-'0')
		}
		if i == len(expression)-1 {
			cur.denominator = v
		}
	}
	last = add(last, cur)

	if last.sign == '-' {
		return fmt.Sprintf("-%v/%v", last.molecular, last.denominator)
	} else {
		return fmt.Sprintf("%v/%v", last.molecular, last.denominator)
	}
}
