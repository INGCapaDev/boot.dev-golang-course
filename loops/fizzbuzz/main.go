package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	FIZZ := "Fizz"
	BUZZ := "Buzz"
	output := make([]string, 0, n)

	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			output = append(output, FIZZ+BUZZ)
		case i%3 == 0:
			output = append(output, FIZZ)
		case i%5 == 0:
			output = append(output, BUZZ)
		default:
			output = append(output, strconv.Itoa(i))
		}
	}
	return output
}

func main() {
	fmt.Println(fizzBuzz(100))
}
