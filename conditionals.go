package main

import "fmt"

func main() {
	messageLen := 30
	maxMessageLen := 20
	fmt.Println("Max message lenght allowed:", maxMessageLen)
	fmt.Println("Trying to send a message of length:", messageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
}
