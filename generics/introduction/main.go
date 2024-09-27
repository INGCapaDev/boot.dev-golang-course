package main

func getLast[T any](s []T) T {
	sLength := len(s)

	if sLength == 0 {
		var zeroV T
		return zeroV
	}

	return s[sLength-1]
}
