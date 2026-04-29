// Complete the getMonthlyPrice function. It accepts a tier (string) as input and returns the monthly price for that tier in pennies. Here are the prices in dollars:
//
// "basic" - $100.00
// "premium" - $150.00
// "enterprise" - $500.00
//
// Convert the prices from dollars to pennies. If the given tier doesn't match any of the above, return 0 pennies.

package main

func getMonthlyPrice(tier string) int {
	var pennies int

	switch tier {
	case "basic":
		pennies = 10000
	case "premium":
		pennies = 15000
	case "enterprise":
		pennies = 50000
	default:
		pennies = 0
	}

	return pennies
}
