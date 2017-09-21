package main

import "fmt"

func fizzBuzz(n int) []string {
	ret := make([]string, n)

	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			ret[i-1] = "FizzBuzz"
		} else if i%5 == 0 {
			ret[i-1] = "Buzz"
		} else if i%3 == 0 {
			ret[i-1] = "Fizz"
		} else {
			ret[i-1] = fmt.Sprintf("%d", i)
		}
	}

	return ret
}
