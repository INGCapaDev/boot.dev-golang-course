package main

import (
	"strconv"
	"strings"
)

func isValidPassword(password string) bool {
	digits := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	passLength := len(password)
	hasUpper, hasDigit := false, false

	if passLength < 5 || passLength > 12 {
		return false
	}

	for i := range password {
		ch := string(password[i])
		isDigit := false

		for _, digit := range digits {
			if strconv.Itoa(digit) == ch {
				hasDigit, isDigit = true, true
			}
		}

		if strings.ToUpper(ch) == string(ch) && !isDigit {
			hasUpper = true
		}
	}

	if !hasDigit || !hasUpper {
		return false
	}

	return true
}
