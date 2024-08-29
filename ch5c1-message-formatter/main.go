package main

import "fmt"

type Formatter interface {
	Format() string
}

type PlainText struct {
	message string
}

type Bold struct {
	message string
}

type Code struct {
	message string
}

func addAroundString(plain string, around string) (formatted string) {
	return fmt.Sprintf("%s%s%s", around, plain, around)
}

func (pt PlainText) Format() string {
	return pt.message
}
func (b Bold) Format() string {
	return addAroundString(b.message, "**")
}
func (c Code) Format() string {
	return addAroundString(c.message, "`")
}

// Don't Touch below this line

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}
