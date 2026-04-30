package main

import "fmt"

func printReports(intro, body, outro string) {
	printCostReport(func(intro string) int {
		return len(intro) * 2
	}, intro)

	printCostReport(func(body string) int {
		return len(body) * 3
	}, body)

	printCostReport(func(outro string) int {
		return len(outro) * 4
	}, outro)
}

// don't touch below this line

func main() {
	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}
