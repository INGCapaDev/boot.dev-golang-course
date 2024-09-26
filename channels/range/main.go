package main

func concurrentFib(n int) []int {
	values := []int{}
	ch := make(chan int, n)

	go fibonacci(n, ch)

	for number := range ch {
		values = append(values, number)
	}
	return values
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
