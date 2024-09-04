package main

import (
	"fmt"
)

func printPrimes(max int) {
	for i := 2; i <= max; i++ {
		prime := true
		if i == 2 {
			fmt.Println(i)
		}
		if i%2 == 0 {
			continue
		}
		for x := 3; x*x <= i; x += 2 {
			if i%x == 0 {
				prime = false
				break
			}
		}
		if prime {
			fmt.Println(i)
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
