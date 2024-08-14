// package main

import "fmt"

func main() {
	//initialize variables here
	var smsSendingLimit int = 10
	var costPerSMS float64 = 10.1
	var hasPermission bool = true
	var username string = "ingcapadev"

	fmt.Printf("%v, %f, %v, %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

}
