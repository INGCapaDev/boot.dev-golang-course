package main

import (
	"fmt"
)

func getSMSErrorString(cost float64, recipient string) string {
	msg := "SMS that costs $%.2f to be sent to '%s' can not be sent"
	return fmt.Sprintf(msg, cost, recipient)
}
