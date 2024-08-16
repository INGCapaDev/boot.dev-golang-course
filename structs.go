// package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

func test(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("============================")
}

func main() {
	test(messageToSend{
		phoneNumber: 12421421412412,
		message:     "Thanks for signing up",
	})
	test(messageToSend{
		phoneNumber: 2343242342,
		message:     "We're so excited to have you",
	})
}
