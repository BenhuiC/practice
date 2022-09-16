package partice

import "fmt"

func fizzBuzz(n int) []string {
	res412 := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			res412 = append(res412, "FizzBuzz")
			continue
		}
		if i%3 == 0 {
			res412 = append(res412, "Fizz")
			continue
		}
		if i%5 == 0 {
			res412 = append(res412, "Buzz")
			continue
		}
		res412 = append(res412, fmt.Sprint(i))
	}
	return res412
}
