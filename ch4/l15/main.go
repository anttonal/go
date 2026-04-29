package main

import "fmt"

func reformat(message string, formatter func(string) string) string {

	var reformattedMessage string

	first := formatter(message)
	second := formatter(first)
	third := formatter(second)

	reformattedMessage = fmt.Sprintf("TEXTIO: %s", third)

	return reformattedMessage
}
