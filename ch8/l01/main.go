package main

func bulkSend(numMessages int) float64 {
	cost := 0.00
	for i := 0; i < numMessages; i++ {
		cost += 1.0 + (float64(i) / 100)
	}
	return cost
}
