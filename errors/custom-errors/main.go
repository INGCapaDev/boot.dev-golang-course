package main

import "fmt"

type divideError struct {
	dividend float64
}

func (err divideError) Error() string {
	sms := "can not divide %v by zero"
	return fmt.Sprintf(sms, err.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}
