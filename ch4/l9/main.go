// One of our clients likes us to send text messages reminding users of life events coming up.

// Fix the bug by adding named return values to the function signature – the bare return at the end is already a naked return that will return them.
// The variables need to be automatically initialized.
// Order them as they appear in the code.
// Do not alter the body of the function.

package main

func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	// don't touch below this line

	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return
}
