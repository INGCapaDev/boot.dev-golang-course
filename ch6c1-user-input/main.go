package main

import (
	"fmt"
)

func validateStatus(status string) error {
	maxCharacters := 140
	statusLength := len(status)
	if statusLength == 0 {
		return fmt.Errorf("status cannot be empty")
	}
	if statusLength > maxCharacters {
		return fmt.Errorf("status exceeds %v characters", maxCharacters)
	}
	return nil
}
