package main

import "fmt"

type email struct {
	isSubscribed bool
	body         string
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

func (e email) cost() int {
	cost := 5
	if e.isSubscribed {
		cost = 2
	}
	return cost * len(e.body)
}

func (e email) format() string {
	message := "Not Subscribed"
	if e.isSubscribed {
		message = "Subscribed"
	}
	return fmt.Sprintf("'%s' | %s", e.body, message)
}

func printExpenseAndFormat(e email) {
	fmt.Printf("%v, %s", e.cost(), e.format())
}

func main() {
	email := email{
		isSubscribed: true,
		body:         "Mock Message",
	}
	printExpenseAndFormat(email)
}
