package main

import "fmt"

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	switch plan {
	case planFree:
		return messages[0:2], nil
	case planPro:
		return messages[:], nil
	default:
		return nil, fmt.Errorf("unsupported plan")
	}

}
